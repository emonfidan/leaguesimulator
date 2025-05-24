import React, { useState, useEffect } from 'react';
import { getPredictions, getLeagueStats } from '../services/api';
import { Prediction, SeasonSimulation } from '../types/types';

const PredictionsPage: React.FC = () => {
  const [predictions, setPredictions] = useState<Prediction[]>([]);
  const [simulation, setSimulation] = useState<SeasonSimulation | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchPredictions = async () => {
    try {
      setLoading(true);
      const data = await getPredictions();
      setPredictions(data.match_predictions);
      setSimulation(data.season_simulation);
    } catch (err) {
      setError('Failed to fetch predictions');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPredictions();
  }, []);

  if (loading) return <div className="text-center py-8">Loading predictions...</div>;
  if (error) return <div className="text-center py-8 text-red-500">{error}</div>;

  return (
    <div className="space-y-8">
      <h1 className="text-2xl font-bold">AI Predictions</h1>
      
      {simulation && (
        <div className="bg-white rounded-lg shadow p-6">
          <h2 className="text-xl font-semibold mb-4">Championship Probabilities</h2>
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            {Object.entries(simulation.championship_probabilities).map(([team, probability]) => (
              <div key={team} className="bg-blue-50 p-4 rounded-lg">
                <h3 className="font-medium">{team}</h3>
                <div className="mt-2">
                  <div className="w-full bg-gray-200 rounded-full h-4">
                    <div 
                      className="bg-blue-600 h-4 rounded-full" 
                      style={{ width: `${probability}%` }}
                    ></div>
                  </div>
                  <p className="text-right mt-1 text-sm font-medium">
                    {probability.toFixed(1)}%
                  </p>
                </div>
              </div>
            ))}
          </div>
          <p className="text-sm text-gray-500 mt-4">
            Based on {simulation.simulation_runs} Monte Carlo simulations
          </p>
        </div>
      )}
      
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-xl font-semibold mb-4">Match Predictions</h2>
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Match</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Prediction</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Confidence</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Win Probabilities</th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {predictions.map((prediction, index) => (
                <tr key={index}>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="font-medium">{prediction.team1} vs {prediction.team2}</div>
                    <div className="text-sm text-gray-500">Expected: {prediction.expected_goals.home.toFixed(1)}-{prediction.expected_goals.away.toFixed(1)}</div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className="font-medium">{prediction.result}</span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="w-full bg-gray-200 rounded-full h-2.5">
                      <div 
                        className="bg-green-600 h-2.5 rounded-full" 
                        style={{ width: `${prediction.confidence}%` }}
                      ></div>
                    </div>
                    <span className="text-sm text-gray-500">{prediction.confidence.toFixed(1)}%</span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex space-x-2">
                      <span className="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">
                        {prediction.team1}: {prediction.win_probabilities.home_win.toFixed(1)}%
                      </span>
                      <span className="text-xs bg-gray-100 text-gray-800 px-2 py-1 rounded">
                        Draw: {prediction.win_probabilities.draw.toFixed(1)}%
                      </span>
                      <span className="text-xs bg-red-100 text-red-800 px-2 py-1 rounded">
                        {prediction.team2}: {prediction.win_probabilities.away_win.toFixed(1)}%
                      </span>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default PredictionsPage;