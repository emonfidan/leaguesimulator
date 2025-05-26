import React, { useState, useEffect } from 'react';
import { getPredictions, initLeague, playNextWeek, getStandings, playAllMatches, resetLeague } from '../services/api';
import LeagueTable from '../components/LeagueTable';
import MatchList from '../components/MatchList';
import WeekSelector from '../components/WeekSelector';
import {  Prediction, SeasonSimulation, StandingsResponse, Match } from '../types/types';

const LeaguePage: React.FC = () => {
  const [standings, setStandings] = useState<StandingsResponse | null>(null);
  const [matches, setMatches] = useState<Match[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [simulation, setSimulation] = useState<SeasonSimulation | null>(null);
  const [needsInitialization, setNeedsInitialization] = useState(false);

  const fetchStandings = async (autoInit = false) => {
    try {
      setLoading(true);
      const data = await getStandings();
      
      // Eğer hiç maç oynanmamışsa (current_week 0 veya standings boş)
      if (!data || data.current_week === 0 || !data.standings || data.standings.length === 0) {
        if (autoInit) {
          // Otomatik olarak league'i initialize et
          await initLeague();
          const newData = await getStandings();
          setStandings(newData);
          setNeedsInitialization(false);
        } else {
          setNeedsInitialization(true);
        }
      } else {
        // Maçlar var, verileri göster
        setStandings(data);
        setNeedsInitialization(false);
      }
    } catch (err) {
      // Eğer standings alınamadıysa, muhtemelen hiç veri yok
      if (autoInit) {
        try {
          await initLeague();
          const newData = await getStandings();
          setStandings(newData);
          setNeedsInitialization(false);
        } catch (initErr) {
          setError('Failed to initialize league');
          console.error(initErr);
        }
      } else {
        setNeedsInitialization(true);
      }
    } finally {
      setLoading(false);
    }
  };

  const fetchPredictions = async () => {
    try {
      const data = await getPredictions();
      setSimulation(data.season_simulation);
    } catch (err) {
      console.error('Failed to fetch predictions:', err);
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
      setNeedsInitialization(true);
    } catch (err) {
      setError('Failed to reset league');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  // Sayfa yüklendiğinde otomatik olarak verileri yükle
  useEffect(() => {
    fetchStandings(true); // autoInit = true
  }, []);

  useEffect(() => {
    if (standings && standings.current_week >= 4) {
      fetchPredictions();
    }
  }, [standings]);

  if (loading) return <div className="text-center py-8">Loading...</div>;
  if (error) return <div className="text-center py-8 text-red-500">{error}</div>;

  return (
    <div className="space-y-8">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">League Management</h1>
        <div className="space-x-2">
          {needsInitialization && (
            <button
              onClick={handleInitLeague}
              className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
            >
              Initialize League
            </button>
          )}
          {standings && !needsInitialization && (
            <>
              <button
                onClick={handlePlayNextWeek}
                disabled={
                  standings?.league_status === 'completed' || !standings?.standings?.length
                }
                className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:bg-gray-400"
              >
                Play Next Week
              </button>

              <button
                onClick={handlePlayAll}
                disabled={
                  standings?.league_status === 'completed' || !standings?.standings?.length
                }
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

      {standings && standings.current_week >= 4 && simulation && (
        <div className="bg-white rounded-lg shadow-sm p-4">
          <h2 className="text-lg font-semibold mb-3">Championship Probabilities</h2>
          <div className="grid grid-cols-1 md:grid-cols-4 gap-2">
            {Object.entries(simulation.championship_probabilities).map(([team, probability]) => (
              <div key={team} className="bg-blue-50 p-2 rounded-md">
                <h3 className="text-sm font-medium">{team}</h3>
                <div className="mt-1">
                  <div className="w-full bg-gray-200 rounded-full h-2.5">
                    <div 
                      className="bg-blue-600 h-2.5 rounded-full" 
                      style={{ width: `${probability}%` }}
                    ></div>
                  </div>
                  <p className="text-right mt-1 text-xs font-medium">
                    {probability.toFixed(1)}%
                  </p>
                </div>
              </div>
            ))}
          </div>
          <p className="text-xs text-gray-500 mt-3">
            Based on {simulation.simulation_runs} Monte Carlo simulations
          </p>
        </div>
      )}

      {standings && !needsInitialization && (
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