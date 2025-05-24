import React from 'react';
import { Team } from '../types/types';

interface LeagueTableProps {
  standings: Team[];
  currentWeek: number;
}

const LeagueTable: React.FC<LeagueTableProps> = ({ standings, currentWeek }) => {
    if (!standings || standings.length === 0) {
        return <div className="p-4">No standings available. Please first reset league.</div>;
    }
    return (
    <div className="overflow-x-auto bg-white rounded-lg shadow">
      <h2 className="text-xl font-bold p-4">League Standings - Week {currentWeek}</h2>
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Position</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Team</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Played</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Won</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Drawn</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Lost</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">GF</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">GA</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">GD</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Points</th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {standings.map((team, index) => (
            <tr key={team.name} className={index < 4 ? 'bg-blue-50' : ''}>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{index + 1}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{team.name}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.played}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.won}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.drawn}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.lost}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.goals_for}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.goals_against}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{team.goal_diff}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-bold text-gray-900">{team.points}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default LeagueTable;