import React, { useState, useEffect } from 'react';
import { initLeague, playNextWeek, getStandings, playAllMatches, resetLeague } from '../services/api';
import LeagueTable from '../components/LeagueTable';
import MatchList from '../components/MatchList';
import WeekSelector from '../components/WeekSelector';
import { StandingsResponse, Match } from '../types/types';

const LeaguePage: React.FC = () => {
  const [standings, setStandings] = useState<StandingsResponse | null>(null);
  const [matches, setMatches] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchStandings = async () => {
    try {
      setLoading(true);
      const data = await getStandings();
      setStandings(data);
    } catch (err) {
      setError('Failed to fetch standings');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleInitLeague = async () => {
    try {
      setLoading(true);
      await initLeague();
      await fetchStandings();
    } catch (err) {
      setError('Failed to initialize league');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handlePlayNextWeek = async () => {
    try {
      setLoading(true);
      const weekData = await playNextWeek();
      setMatches(weekData.matches);
      await fetchStandings();
    } catch (err) {
      setError('Failed to play next week');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handlePlayAll = async () => {
    try {
      setLoading(true);
      await playAllMatches();
      await fetchStandings();
    } catch (err) {
      setError('Failed to play all matches');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleReset = async () => {
    try {
      setLoading(true);
      await resetLeague();
      setStandings(null);
      setMatches([]);
    } catch (err) {
      setError('Failed to reset league');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStandings();
  }, []);

  if (loading) return <div className="text-center py-8">Loading...</div>;
  if (error) return <div className="text-center py-8 text-red-500">{error}</div>;

  return (
    <div className="space-y-8">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">League Management</h1>
        <div className="space-x-2">
          {!standings && (
            <button
              onClick={handleInitLeague}
              className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
            >
              Initialize League
            </button>
          )}
          {standings && (
            <>
              <button
                onClick={handlePlayNextWeek}
                disabled={standings?.league_status === 'completed'}
                className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:bg-gray-400"
              >
                Play Next Week
              </button>
              <button
                onClick={handlePlayAll}
                disabled={standings?.league_status === 'completed'}
                className="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700 disabled:bg-gray-400"
              >
                Play All Remaining
              </button>
              <button
                onClick={handleReset}
                className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
              >
                Reset League
              </button>
            </>
          )}
        </div>
      </div>

      {standings && (
        <>
          <LeagueTable standings={standings.standings} currentWeek={standings.current_week} />
          <WeekSelector currentWeek={standings.current_week} />
          <MatchList matches={matches} />
        </>
      )}
    </div>
  );
};

export default LeaguePage;