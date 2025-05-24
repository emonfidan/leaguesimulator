import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { getTeamAnalysis, getHeadToHead } from '../services/api';
import { TeamAnalysis } from '../types/types';

const TeamPage: React.FC = () => {
  const { teamName } = useParams<{ teamName: string }>();
  const [analysis, setAnalysis] = useState<TeamAnalysis | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchTeamAnalysis = async () => {
    if (!teamName) return;
    
    try {
      setLoading(true);
      const data = await getTeamAnalysis(teamName);
      setAnalysis(data);
    } catch (err) {
      setError('Failed to fetch team analysis');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTeamAnalysis();
  }, [teamName]);

  if (loading) return <div className="text-center py-8">Loading team analysis...</div>;
  if (error) return <div className="text-center py-8 text-red-500">{error}</div>;
  if (!analysis) return <div className="text-center py-8">No data available</div>;

  return (
    <div className="space-y-8">
      <h1 className="text-2xl font-bold">{teamName} Analysis</h1>
      
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-xl font-semibold mb-4">Performance</h2>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <h3 className="font-medium mb-2">Home Performance</h3>
            <div className="space-y-2">
              <p>Wins: {analysis.performance.home.wins}</p>
              <p>Draws: {analysis.performance.home.draws}</p>
              <p>Losses: {analysis.performance.home.losses}</p>
            </div>
          </div>
          
          <div>
            <h3 className="font-medium mb-2">Away Performance</h3>
            <div className="space-y-2">
              <p>Wins: {analysis.performance.away.wins}</p>
              <p>Draws: {analysis.performance.away.draws}</p>
              <p>Losses: {analysis.performance.away.losses}</p>
            </div>
          </div>
        </div>
      </div>
      
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-xl font-semibold mb-4">Goal Statistics</h2>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <h3 className="font-medium mb-2">Goals For</h3>
            <p>Average per match: {analysis.goals.for.average.toFixed(1)}</p>
            <div className="mt-4">
              <h4 className="text-sm font-medium mb-2">Distribution</h4>
              <ul className="space-y-1">
                {Object.entries(analysis.goals.for.distribution).map(([goals, percentage]) => (
                  <li key={goals} className="flex items-center">
                    <span className="w-16">{goals} goals:</span>
                    <div className="flex-1 bg-gray-200 rounded-full h-2.5">
                      <div 
                        className="bg-green-600 h-2.5 rounded-full" 
                        style={{ width: `${percentage}%` }}
                      ></div>
                    </div>
                    <span className="ml-2 text-sm w-12">{percentage}%</span>
                  </li>
                ))}
              </ul>
            </div>
          </div>
          
          <div>
            <h3 className="font-medium mb-2">Goals Against</h3>
            <p>Average per match: {analysis.goals.against.average.toFixed(1)}</p>
            <div className="mt-4">
              <h4 className="text-sm font-medium mb-2">Distribution</h4>
              <ul className="space-y-1">
                {Object.entries(analysis.goals.against.distribution).map(([goals, percentage]) => (
                  <li key={goals} className="flex items-center">
                    <span className="w-16">{goals} goals:</span>
                    <div className="flex-1 bg-gray-200 rounded-full h-2.5">
                      <div 
                        className="bg-red-600 h-2.5 rounded-full" 
                        style={{ width: `${percentage}%` }}
                      ></div>
                    </div>
                    <span className="ml-2 text-sm w-12">{percentage}%</span>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </div>
      
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-xl font-semibold mb-4">Recent Form</h2>
        <div className="flex space-x-2">
          {analysis.form.map((result, index) => (
            <span 
              key={index} 
              className={`w-8 h-8 rounded-full flex items-center justify-center text-white ${
                result === 'W' ? 'bg-green-500' :
                result === 'D' ? 'bg-yellow-500' : 'bg-red-500'
              }`}
            >
              {result}
            </span>
          ))}
        </div>
      </div>
    </div>
  );
};

export default TeamPage;