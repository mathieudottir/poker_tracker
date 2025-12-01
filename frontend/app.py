import streamlit as st
import streamlit.components.v1 as components
import requests
import pandas as pd
import plotly.graph_objects as go
import plotly.express as px
from datetime import datetime
import os
import base64

# Backend API URL
API_URL = os.getenv("API_URL", "http://localhost:8080/api")

# Page configuration
st.set_page_config(
    page_title="Winamax Expresso Tracker",
    page_icon="🃏",
    layout="wide",
    initial_sidebar_state="expanded"
)

# Initialize session state
if 'user' not in st.session_state:
    st.session_state.user = None
if 'user_id' not in st.session_state:
    st.session_state.user_id = None

def api_call(endpoint, method="GET", data=None):
    """Make API call to backend"""
    url = f"{API_URL}{endpoint}"
    try:
        if method == "GET":
            response = requests.get(url)
        elif method == "POST":
            response = requests.post(url, json=data)
        elif method == "PUT":
            response = requests.put(url, json=data)
        elif method == "DELETE":
            response = requests.delete(url)

        if response.status_code in [200, 201]:
            return response.json()
        else:
            st.error(f"API Error: {response.status_code} - {response.text}")
            return None
    except Exception as e:
        st.error(f"Connection Error: {str(e)}")
        return None

def login_page():
    """Login page"""
    st.title("🃏 Winamax Expresso Tracker")
    st.subheader("Login")

    # Dev mode login button
    col1, col2 = st.columns([3, 1])
    with col1:
        username = st.text_input("Username")
        password = st.text_input("Password", type="password")

        if st.button("Login", type="primary"):
            result = api_call("/auth/login", "POST", {
                "username": username,
                "password": password
            })
            if result:
                st.session_state.user = result
                st.session_state.user_id = result['id']
                st.rerun()

    with col2:
        st.write("")
        st.write("")
        if st.button("🔧 Login as Mathieu (DEV)"):
            result = api_call("/auth/dev-login", "POST", {
                "username": "mathieu"
            })
            if result:
                st.session_state.user = result
                st.session_state.user_id = result['id']
                st.rerun()

    st.divider()

    # Registration
    with st.expander("Register New Account"):
        reg_username = st.text_input("Username", key="reg_user")
        reg_password = st.text_input("Password", type="password", key="reg_pass")
        reg_player = st.text_input("Winamax Player Name", key="reg_player")

        if st.button("Register"):
            result = api_call("/auth/register", "POST", {
                "username": reg_username,
                "password": reg_password,
                "player_name": reg_player
            })
            if result:
                st.success("Registration successful! Please login.")

def import_page():
    """Import page"""
    st.title("📥 Import Hand Histories")

    user_id = st.session_state.user_id

    # Watcher status (collapsed by default)
    with st.expander("⚙️ Surveillance Automatique (Watcher) - Avancé"):
        st.warning("""
        ⚠️ **Ce mode est pour utilisateurs avancés seulement.**

        Le watcher surveille un **dossier sur le SERVEUR** (pas votre PC).
        Pour la plupart des utilisateurs, l'upload de fichiers ci-dessous est bien plus simple !
        """)

        watcher_status = api_call(f"/watcher/status/{user_id}")

        col1, col2 = st.columns([2, 1])
        with col1:
            if watcher_status and watcher_status.get('running'):
                st.success(f"✅ Watcher actif sur: {watcher_status.get('directory')}")
            else:
                st.info("⚠️ Watcher inactif")

        with col2:
            if watcher_status and watcher_status.get('running'):
                if st.button("⏸️ Arrêter", type="secondary"):
                    result = api_call("/watcher/stop", "POST", {"user_id": user_id})
                    if result:
                        st.success("Watcher arrêté")
                    st.rerun()
            else:
                if st.button("▶️ Démarrer", type="primary"):
                    user_data = api_call(f"/user/{user_id}")
                    if user_data and user_data.get('hh_directory'):
                        result = api_call("/watcher/start", "POST", {
                            "user_id": user_id,
                            "directory": user_data['hh_directory']
                        })
                        if result:
                            st.success("Watcher démarré!")
                            st.rerun()
                        else:
                            st.error("❌ Le dossier n'existe pas sur le serveur. Allez dans Settings pour le créer.")
                    else:
                        st.error("⚠️ Configurez d'abord le chemin dans Settings")

    st.divider()

    # File upload import
    st.subheader("📤 Importer vos Hand Histories")

    with st.expander("📖 **Comment importer tous vos fichiers en 3 clics**", expanded=True):
        st.markdown("""
        1. 📁 Sur votre PC, ouvrez votre dossier **Winamax HandHistory**
           - Windows: `C:\\Users\\VotreNom\\AppData\\Local\\Programs\\WinamaxPoker\\HandHistory\\VotreNom`
           - Mac: `~/Library/Application Support/WinamaxPoker/HandHistory/VotreNom`

        2. ⌨️ Sélectionnez **TOUS** les fichiers .txt
           - Windows: **Ctrl + A**
           - Mac: **Cmd + A**

        3. 🎯 **Glissez-déposez** les fichiers ci-dessous
           - OU cliquez "Browse files" et validez

        💡 **Vous pouvez importer des milliers de fichiers d'un coup !**
        """)

    uploaded_files = st.file_uploader(
        "📁 Glissez-déposez vos fichiers .txt ici",
        type=['txt'],
        accept_multiple_files=True,
        help="Faites Ctrl+A dans votre dossier HandHistory pour tout sélectionner, puis glissez ici",
        key="file_uploader"
    )

    if uploaded_files:
        st.success(f"✅ **{len(uploaded_files)} fichier(s)** prêts à être importés")

        if st.button("🚀 IMPORTER", type="primary", use_container_width=True):
            with st.spinner(f"⏳ Traitement de {len(uploaded_files)} fichier(s)..."):
                import tempfile
                import shutil

                # Create temp directory
                temp_dir = tempfile.mkdtemp()

                try:
                    # Save uploaded files with progress
                    progress_text = st.empty()
                    progress_bar = st.progress(0)

                    for idx, uploaded_file in enumerate(uploaded_files):
                        progress_text.text(f"📥 Téléchargement {idx+1}/{len(uploaded_files)}: {uploaded_file.name}")
                        file_path = os.path.join(temp_dir, uploaded_file.name)
                        with open(file_path, 'wb') as f:
                            f.write(uploaded_file.getbuffer())
                        progress_bar.progress((idx + 1) / len(uploaded_files))

                    progress_text.text("🔍 Analyse et import des tournois...")

                    # Import from temp directory
                    result = api_call("/import/directory", "POST", {
                        "user_id": user_id,
                        "directory": temp_dir
                    })

                    progress_bar.empty()
                    progress_text.empty()

                    if result:
                        st.success(f"""
                        🎉 **Import réussi !**

                        - 📊 **{result.get('count', 0)} tournois** importés
                        - 📁 Depuis **{len(uploaded_files)} fichiers**
                        """)
                        st.balloons()
                        st.rerun()
                    else:
                        st.error("❌ Erreur lors de l'import des tournois")
                except Exception as e:
                    st.error(f"❌ Erreur: {str(e)}")
                finally:
                    # Cleanup temp directory
                    shutil.rmtree(temp_dir, ignore_errors=True)

    st.divider()

    # Manual import
    st.subheader("📁 Import from Server Directory")

    import_dir = st.text_input("Directory Path", value="./examples")

    if st.button("Import from Directory"):
        with st.spinner("Importing tournaments..."):
            result = api_call("/import/directory", "POST", {
                "user_id": user_id,
                "directory": import_dir
            })
            if result:
                st.success(f"✅ Imported {result.get('count', 0)} tournaments")
                st.rerun()

    st.divider()

    # Import logs
    st.subheader("Import Logs")
    logs = api_call("/import/logs")
    if logs:
        for log in logs[-20:]:  # Show last 20 logs
            st.text(log)

def results_page():
    """Results page"""
    st.title("📊 My Results")

    user_id = st.session_state.user_id

    # Get tournaments
    tournaments = api_call(f"/tournaments/{user_id}")

    if not tournaments:
        st.warning("No tournaments found. Import some hand histories first!")
        return

    # Convert to DataFrame
    df = pd.DataFrame(tournaments)
    df['start_time'] = pd.to_datetime(df['start_time'])
    df['buyin_euros'] = df['buyin_cents'] / 100
    df['net_result_euros'] = df['net_result_cents'] / 100
    df['ev_euros'] = df['ev_cents'] / 100

    # Date range filter
    st.subheader("🗓️ Filtres")
    col1, col2 = st.columns(2)
    with col1:
        min_date = df['start_time'].min().date()
        max_date = df['start_time'].max().date()
        start_date = st.date_input("Date de début", value=min_date, min_value=min_date, max_value=max_date)
    with col2:
        end_date = st.date_input("Date de fin", value=max_date, min_value=min_date, max_value=max_date)

    # Filter dataframe
    mask = (df['start_time'].dt.date >= start_date) & (df['start_time'].dt.date <= end_date)
    df = df[mask]

    if len(df) == 0:
        st.warning("Aucun tournoi dans cette plage de dates")
        return

    # Summary metrics
    col1, col2, col3, col4 = st.columns(4)

    with col1:
        st.metric("Total Tournaments", len(df))
    with col2:
        total_profit = df['net_result_euros'].sum()
        st.metric("Total Profit", f"€{total_profit:.2f}")
    with col3:
        avg_roi = (df['net_result_euros'].sum() / (df['buyin_euros'].sum() + df['rake_cents'].sum()/100)) * 100
        st.metric("ROI", f"{avg_roi:.2f}%")
    with col4:
        total_ev = df['ev_euros'].sum()
        st.metric("Total EV", f"€{total_ev:.2f}")

    st.divider()

    # Bankroll curve
    st.subheader("Bankroll Evolution")
    df_sorted = df.sort_values('start_time').reset_index(drop=True)
    df_sorted['cumulative'] = df_sorted['net_result_euros'].cumsum()
    df_sorted['tournament_number'] = range(1, len(df_sorted) + 1)

    fig = go.Figure()
    fig.add_trace(go.Scatter(
        x=df_sorted['tournament_number'],
        y=df_sorted['cumulative'],
        mode='lines+markers',
        name='Bankroll',
        line=dict(color='green', width=2),
        hovertemplate='Tournoi #%{x}<br>Profit: €%{y:.2f}<br>Date: %{customdata}<extra></extra>',
        customdata=df_sorted['start_time'].dt.strftime('%Y-%m-%d %H:%M')
    ))
    fig.update_layout(
        title="Évolution de la Bankroll",
        xaxis_title="Numéro de Tournoi",
        yaxis_title="Profit (€)",
        hovermode='x unified'
    )
    st.plotly_chart(fig, use_container_width=True)

    # EV curve
    st.subheader("Évolution EV vs Résultats Réels")
    df_sorted['cumulative_ev'] = df_sorted['ev_euros'].cumsum()

    fig2 = go.Figure()
    fig2.add_trace(go.Scatter(
        x=df_sorted['tournament_number'],
        y=df_sorted['cumulative_ev'],
        mode='lines+markers',
        name='EV',
        line=dict(color='blue', width=2),
        hovertemplate='Tournoi #%{x}<br>EV: €%{y:.2f}<extra></extra>'
    ))
    fig2.add_trace(go.Scatter(
        x=df_sorted['tournament_number'],
        y=df_sorted['cumulative'],
        mode='lines+markers',
        name='Résultats Réels',
        line=dict(color='green', width=2, dash='dash'),
        hovertemplate='Tournoi #%{x}<br>Profit: €%{y:.2f}<extra></extra>'
    ))
    fig2.update_layout(
        title="EV vs Résultats Réels",
        xaxis_title="Numéro de Tournoi",
        yaxis_title="Valeur (€)",
        hovermode='x unified'
    )
    st.plotly_chart(fig2, use_container_width=True)

    # Multiplier distribution
    st.subheader("Distribution des Multiplicateurs")
    try:
        mult_dist = api_call(f"/stats/{user_id}/multipliers")
        if mult_dist and isinstance(mult_dist, dict):
            mult_df = pd.DataFrame(list(mult_dist.items()), columns=['Multiplicateur', 'Nombre'])
            mult_df['Multiplicateur'] = mult_df['Multiplicateur'].astype(int)
            mult_df = mult_df.sort_values('Multiplicateur')
            fig3 = px.bar(mult_df, x='Multiplicateur', y='Nombre',
                         title='Multiplicateurs Obtenus',
                         color='Nombre',
                         color_continuous_scale='Greens')
            st.plotly_chart(fig3, use_container_width=True)
        else:
            # Calculate from filtered data if API fails
            mult_counts = df['multiplier'].value_counts().sort_index()
            fig3 = px.bar(x=mult_counts.index, y=mult_counts.values,
                         labels={'x': 'Multiplicateur', 'y': 'Nombre'},
                         title='Multiplicateurs Obtenus')
            st.plotly_chart(fig3, use_container_width=True)
    except Exception as e:
        st.error(f"Erreur lors du chargement des multiplicateurs: {str(e)}")

    st.divider()

    # Tournament table
    st.subheader("Tournament History")
    display_df = df[['start_time', 'buyin_euros', 'multiplier', 'hero_rank', 'net_result_euros', 'ev_euros']].copy()
    display_df.columns = ['Date', 'Buy-in (€)', 'Multiplier', 'Finish', 'Profit (€)', 'EV (€)']
    st.dataframe(display_df, use_container_width=True)

def statistics_page():
    """Statistics page"""
    st.title("📈 Statistics")

    user_id = st.session_state.user_id

    # Get stats
    stats = api_call(f"/stats/{user_id}")

    if not stats:
        st.warning("No statistics available yet.")
        return

    # Overall stats
    st.subheader("Overall Statistics")

    col1, col2, col3 = st.columns(3)
    with col1:
        st.metric("Total Tournaments", stats.get('total_tournaments', 0))
        st.metric("Total Hands", stats.get('total_hands', 0))
    with col2:
        st.metric("Net Profit", f"€{stats.get('net_result_cents', 0)/100:.2f}")
        st.metric("EV", f"€{stats.get('ev_cents', 0)/100:.2f}")
    with col3:
        st.metric("Net ROI", f"{stats.get('net_roi', 0):.2f}%")
        st.metric("EV ROI", f"{stats.get('ev_roi', 0):.2f}%")

    st.divider()

    # Rake stats
    st.subheader("Rake & Rakeback")

    col1, col2 = st.columns(2)
    with col1:
        st.metric("Total Rake Paid", f"€{stats.get('total_rake_cents', 0)/100:.2f}")
    with col2:
        st.metric("Rakeback Earned", f"€{stats.get('total_rakeback_cents', 0)/100:.2f}")

    st.divider()

    # Results by buy-in
    st.subheader("Results by Buy-in")
    by_buyin = api_call(f"/stats/{user_id}/bybuyin")

    if by_buyin:
        buyin_data = []
        for buyin_cents, data in by_buyin.items():
            buyin_data.append({
                'Buy-in (€)': int(buyin_cents) / 100,
                'Tournaments': data.get('total_tournaments', 0),
                'Profit (€)': data.get('net_result_cents', 0) / 100,
                'EV (€)': data.get('ev_cents', 0) / 100,
                'ROI (%)': data.get('net_roi', 0)
            })

        buyin_df = pd.DataFrame(buyin_data)
        st.dataframe(buyin_df, use_container_width=True)

def settings_page():
    """Settings page"""
    st.title("⚙️ Settings")

    user_id = st.session_state.user_id
    user = api_call(f"/user/{user_id}")

    if not user:
        st.error("Failed to load user settings")
        return

    st.subheader("User Settings")

    player_name = st.text_input("Winamax Player Name", value=user.get('player_name', ''))

    st.subheader("📁 Hand History Directory (pour Watcher)")
    st.info("""
    ⚠️ **Important:** Ce chemin est sur le SERVEUR, pas sur votre PC.

    - ✅ **Pour importer depuis votre PC** → Utilisez l'upload de fichiers dans l'onglet Import
    - ⚙️ **Pour surveiller un dossier serveur** → Configurez ce chemin (avancé)
    """)

    hh_directory = st.text_input(
        "Chemin serveur (optionnel)",
        value=user.get('hh_directory', ''),
        placeholder="/home/mathieudottir/winamax_hh",
        help="Chemin absolu sur le serveur pour la surveillance automatique"
    )

    if st.button("📁 Créer le dossier sur le serveur"):
        import os
        server_path = "/home/mathieudottir/winamax_hh"
        try:
            os.makedirs(server_path, exist_ok=True)
            st.success(f"✅ Dossier créé: {server_path}")
            hh_directory = server_path
        except Exception as e:
            st.error(f"❌ Erreur: {str(e)}")

    # Winamax status
    rakeback_statuses = api_call("/winamax/rakeback")
    status_names = [s['status_name'] for s in rakeback_statuses] if rakeback_statuses else []

    current_status = user.get('wina_status', 'Aluminium')
    status_index = status_names.index(current_status) if current_status in status_names else 0

    wina_status = st.selectbox(
        "Winamax Status",
        status_names,
        index=status_index
    )

    dev_mode = st.checkbox("Dev Mode", value=user.get('dev_mode', False))

    if st.button("Save Settings", type="primary"):
        result = api_call(f"/user/{user_id}/settings", "PUT", {
            "player_name": player_name,
            "wina_status": wina_status,
            "hh_directory": hh_directory,
            "dev_mode": dev_mode
        })
        if result:
            st.success("Settings saved successfully!")
            st.rerun()

    st.divider()

    # Database reset
    st.subheader("⚠️ Danger Zone")
    st.warning("The following action will delete ALL your data permanently!")

    if 'confirm_reset' not in st.session_state:
        st.session_state.confirm_reset = False

    if not st.session_state.confirm_reset:
        if st.button("🗑️ Reset My Database", type="secondary"):
            st.session_state.confirm_reset = True
            st.rerun()
    else:
        st.error("⚠️ ATTENTION: Cette action est irréversible!")
        col1, col2 = st.columns(2)
        with col1:
            if st.button("✅ OUI, supprimer toutes mes données", type="primary"):
                api_call(f"/user/{user_id}/data", "DELETE")
                st.success("Toutes les données supprimées. Déconnexion...")
                st.session_state.user = None
                st.session_state.user_id = None
                st.session_state.confirm_reset = False
                st.rerun()
        with col2:
            if st.button("❌ Annuler"):
                st.session_state.confirm_reset = False
                st.rerun()

def main():
    """Main app"""

    # Check if logged in
    if not st.session_state.user:
        login_page()
        return

    # Sidebar
    with st.sidebar:
        st.title("Navigation")
        user = st.session_state.user
        st.write(f"👤 {user.get('player_name', user.get('username'))}")

        page = st.radio(
            "Go to",
            ["Import", "My Results", "Statistics", "Settings"],
            label_visibility="collapsed"
        )

        st.divider()

        if st.button("Logout"):
            st.session_state.user = None
            st.session_state.user_id = None
            st.rerun()

    # Show selected page
    if page == "Import":
        import_page()
    elif page == "My Results":
        results_page()
    elif page == "Statistics":
        statistics_page()
    elif page == "Settings":
        settings_page()

if __name__ == "__main__":
    main()
