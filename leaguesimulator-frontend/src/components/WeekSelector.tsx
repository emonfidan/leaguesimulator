import React from 'react';

interface WeekSelectorProps {
  currentWeek: number;
}

const WeekSelector: React.FC<WeekSelectorProps> = ({ currentWeek }) => {
  const totalWeeks = 3; // Based on your league structure
  
  return (
    <div className="bg-white rounded-lg shadow p-4">
      <h3 className="font-medium mb-2">Season Progress</h3>
      <div className="flex items-center">
        <div className="flex-1 bg-gray-200 rounded-full h-2.5">
          <div 
            className="bg-blue-600 h-2.5 rounded-full" 
            style={{ width: `${(currentWeek / totalWeeks) * 100}%` }}
          ></div>
        </div>
        <span className="ml-4 text-sm font-medium">
          Week {currentWeek} of {totalWeeks}
        </span>
      </div>
    </div>
  );
};

export default WeekSelector;