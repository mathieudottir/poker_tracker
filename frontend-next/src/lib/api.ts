const API_URL = process.env.NEXT_PUBLIC_API_URL || (typeof window !== 'undefined' ? `http://${window.location.hostname}:8080/api` : 'http://localhost:8080/api');

export interface Tournament {
  id: number;
  user_id: number;
  buyin_cents: number;
  rake_cents: number;
  multiplier: number;
  prize1_cents: number;
  prize2_cents: number;
  prize3_cents: number;
  hero_rank: number;
  start_time: string;
  end_time: string;
  hands_count: number;
  net_result_cents: number;
  ev_cents: number;
  chips_ev: number;
  import_file: string;
  tournament_id: string;
}

export interface User {
  id: number;
  username: string;
  player_name: string;
  wina_status: string;
  hh_directory: string;
  dev_mode: boolean;
}

export interface Stats {
  total_tournaments: number;
  net_result_cents: number;
  ev_cents: number;
  total_rake_cents: number;
  total_rakeback: number;
  net_roi: number;
  ev_roi: number;
  winrate: number;
}

async function apiCall<T>(
  endpoint: string,
  method: 'GET' | 'POST' | 'PUT' | 'DELETE' = 'GET',
  data?: any
): Promise<T> {
  const options: RequestInit = {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
  };

  if (data && method !== 'GET') {
    options.body = JSON.stringify(data);
  }

  const response = await fetch(`${API_URL}${endpoint}`, options);

  if (!response.ok) {
    throw new Error(`API Error: ${response.status} - ${response.statusText}`);
  }

  return response.json();
}

export const api = {
  // Auth
  login: (username: string, password: string) =>
    apiCall<User>('/auth/login', 'POST', { username, password }),

  devLogin: (username: string) =>
    apiCall<User>('/auth/dev-login', 'POST', { username }),

  register: (username: string, password: string, playerName: string) =>
    apiCall<User>('/auth/register', 'POST', {
      username,
      password,
      player_name: playerName,
    }),

  // User
  getUser: (userId: number) => apiCall<User>(`/user/${userId}`),

  updateSettings: (userId: number, playerName: string, winaStatus: string, hhDirectory: string, devMode: boolean) =>
    apiCall<{ status: string }>(`/user/${userId}/settings`, 'PUT', {
      player_name: playerName,
      wina_status: winaStatus,
      hh_directory: hhDirectory,
      dev_mode: devMode,
    }),

  deleteUserData: (userId: number) =>
    apiCall<{ status: string }>(`/user/${userId}/data`, 'DELETE'),

  // Tournaments
  getTournaments: (userId: number) =>
    apiCall<Tournament[]>(`/tournaments/${userId}`),

  // Stats
  getStats: (userId: number) => apiCall<Stats>(`/stats/${userId}`),

  getMultipliers: (userId: number) =>
    apiCall<Record<string, number>>(`/stats/${userId}/multipliers`),

  // Import
  importDirectory: (userId: number, directory: string) =>
    apiCall<{ count: number }>('/import/directory', 'POST', {
      user_id: userId,
      directory,
    }),

  importFiles: async (userId: number, files: File[]) => {
    const formData = new FormData();
    files.forEach((file) => {
      formData.append('files', file);
    });
    formData.append('user_id', userId.toString());

    const response = await fetch(`${API_URL}/import/files`, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Upload failed (${response.status}): ${errorText || response.statusText}`);
    }

    return response.json();
  },
};
