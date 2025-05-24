import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const initLeague = async () => {
  const response = await axios.post(`${API_BASE_URL}/init-league`);
  return response.data;
};

export const playNextWeek = async () => {
  const response = await axios.post(`${API_BASE_URL}/next-week`);
  return response.data;
};

export const getStandings = async () => {
  const response = await axios.get(`${API_BASE_URL}/standings`);
  return response.data;
};

export const getPredictions = async () => {
  const response = await axios.get(`${API_BASE_URL}/predict`);
  return response.data;
};

export const editMatchResult = async (data: {
  week: number;
  team1: string;
  team2: string;
  score1: number;
  score2: number;
  reason?: string;
}) => {
  const response = await axios.post(`${API_BASE_URL}/edit-result`, data);
  return response.data;
};

export const playAllMatches = async () => {
  const response = await axios.post(`${API_BASE_URL}/play-all`);
  return response.data;
};

export const getTeamAnalysis = async (teamName: string) => {
  const response = await axios.get(`${API_BASE_URL}/team/${teamName}/analysis`);
  return response.data;
};

export const getHeadToHead = async (team1: string, team2: string) => {
  const response = await axios.get(`${API_BASE_URL}/head-to-head/${team1}/${team2}`);
  return response.data;
};

export const getLeagueStats = async () => {
  const response = await axios.get(`${API_BASE_URL}/league-stats`);
  return response.data;
};

export const resetLeague = async () => {
  const response = await axios.post(`${API_BASE_URL}/reset`);
  return response.data;
};