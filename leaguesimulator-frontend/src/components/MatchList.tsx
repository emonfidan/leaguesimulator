import React, { useState } from 'react';
import { Match } from '../types/types';
import { editMatchResult } from '../services/api';

interface MatchListProps {
  matches: Match[];
  onMatchEdited?: () => void; // Callback to refresh standings after edit
  selectedWeek?: number | null;
}

const MatchList: React.FC<MatchListProps> = ({ matches, onMatchEdited }) => {
  const [editingMatch, setEditingMatch] = useState<number | null>(null);
  const [editScore1, setEditScore1] = useState<number>(0);
  const [editScore2, setEditScore2] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);

  if (matches.length === 0) return null;

  // Show edit button only for the 2 most recent matches
  const recentMatches = matches.slice(0, 2);
  const olderMatches = matches.slice(2);

  const handleEditClick = (index: number, match: Match) => {
    setEditingMatch(index);
    setEditScore1(match.score1);
    setEditScore2(match.score2);
  };

  const handleSaveEdit = async (match: Match) => {
    try {
      setLoading(true);
      console.log('Sending edit request:', {
        week: match.week,
        team1: match.team1,
        team2: match.team2,
        score1: editScore1,
        score2: editScore2
      });
      
      await editMatchResult({
        week: match.week,
        team1: match.team1,
        team2: match.team2,
        score1: editScore1,
        score2: editScore2
      });
      
      setEditingMatch(null);
      if (onMatchEdited) {
        onMatchEdited(); // Refresh standings
      }
    } catch (error) {
      console.error('Failed to edit match result:', error);
      if (error instanceof Error) {
        console.error('Error details:', error.message);
        alert(`Failed to edit match result: ${error.message}`);
      } else if (error && typeof error === 'object' && 'response' in error) {
        const axiosError = error as any;
        console.error('Error details:', axiosError.response?.data);
        console.error('Request data was:', {
          week: match.week,
          team1: match.team1,
          team2: match.team2,
          score1: editScore1,
          score2: editScore2
        });
        alert(`Failed to edit match result: ${axiosError.response?.data?.message || JSON.stringify(axiosError.response?.data) || 'Unknown error'}`);
      } else {
        alert('Failed to edit match result: Unknown error');
      }
    } finally {
      setLoading(false);
    }
  };

  const handleCancelEdit = () => {
    setEditingMatch(null);
  };

  const renderMatch = (match: Match, index: number, showEditButton: boolean) => (
    <div key={index} className="border-b border-gray-200 pb-4 last:border-0 last:pb-0">
      <div className="flex justify-between items-center">
        <div className="text-lg font-medium">
          <span className={match.score1 > match.score2 ? 'font-bold' : ''}>{match.team1}</span>
          {' '}vs{' '}
          <span className={match.score2 > match.score1 ? 'font-bold' : ''}>{match.team2}</span>
        </div>
        
        <div className="flex items-center space-x-4">
          {editingMatch === index ? (
            <div className="flex items-center space-x-2">
              <input
                type="number"
                value={editScore1}
                onChange={(e) => setEditScore1(parseInt(e.target.value) || 0)}
                className="w-16 px-2 py-1 border rounded text-center"
                min="0"
              />
              <span>-</span>
              <input
                type="number"
                value={editScore2}
                onChange={(e) => setEditScore2(parseInt(e.target.value) || 0)}
                className="w-16 px-2 py-1 border rounded text-center"
                min="0"
              />
              <button
                onClick={() => handleSaveEdit(match)}
                disabled={loading}
                className="bg-green-500 text-white px-2 py-1 rounded text-sm hover:bg-green-600 disabled:bg-gray-400"
              >
                Save
              </button>
              <button
                onClick={handleCancelEdit}
                disabled={loading}
                className="bg-gray-500 text-white px-2 py-1 rounded text-sm hover:bg-gray-600"
              >
                Cancel
              </button>
            </div>
          ) : (
            <div className="flex items-center space-x-2">
              <div className="text-xl font-bold">
                {match.score1} - {match.score2}
              </div>
              {showEditButton && (
                <button
                  onClick={() => handleEditClick(index, match)}
                  className="bg-blue-500 text-white px-2 py-1 rounded text-sm hover:bg-blue-600"
                >
                  Edit
                </button>
              )}
            </div>
          )}
        </div>
      </div>
      <div className="text-sm text-gray-500 mt-1">
        Week {match.week} • {match.score1 > match.score2 ? `${match.team1} wins` : 
          match.score2 > match.score1 ? `${match.team2} wins` : 'Draw'}
      </div>
    </div>
  );

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h2 className="text-xl font-semibold mb-4">Recent Matches</h2>
      <div className="space-y-4">
        {recentMatches.map((match, index) => renderMatch(match, index, true))}
        {olderMatches.map((match, index) => renderMatch(match, index + 2, false))}
      </div>
    </div>
  );
};

export default MatchList;