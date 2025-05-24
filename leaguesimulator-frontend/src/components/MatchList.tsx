import React from 'react';
import { Match } from '../types/types';

interface MatchListProps {
  matches: Match[];
}

const MatchList: React.FC<MatchListProps> = ({ matches }) => {
  if (matches.length === 0) return null;

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h2 className="text-xl font-semibold mb-4">Recent Matches</h2>
      <div className="space-y-4">
        {matches.map((match, index) => (
          <div key={index} className="border-b border-gray-200 pb-4 last:border-0 last:pb-0">
            <div className="flex justify-between items-center">
              <div className="text-lg font-medium">
                <span className={match.score1 > match.score2 ? 'font-bold' : ''}>{match.team1}</span>
                {' '}vs{' '}
                <span className={match.score2 > match.score1 ? 'font-bold' : ''}>{match.team2}</span>
              </div>
              <div className="text-xl font-bold">
                {match.score1} - {match.score2}
              </div>
            </div>
            <div className="text-sm text-gray-500 mt-1">
              Week {match.week} • {match.score1 > match.score2 ? `${match.team1} wins` : 
                match.score2 > match.score1 ? `${match.team2} wins` : 'Draw'}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default MatchList;