import React from 'react';
import { Link } from 'react-router-dom';

const Home: React.FC = () => {
  return (
    <div className="space-y-8">
      <div className="bg-white rounded-lg shadow p-6">
        <h1 className="text-3xl font-bold mb-4">Football League Simulator</h1>
        <p className="text-lg mb-6">
          Simulate a football league with AI-powered predictions and detailed analytics.
          Manage teams, track standings, and get insights into match outcomes.
        </p>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="bg-blue-50 p-4 rounded-lg border border-blue-100">
            <h2 className="text-xl font-semibold mb-2 text-blue-800">League Management</h2>
            <p className="mb-4">Initialize the league, play matches week by week, and track team standings.</p>
            <Link to="/league" className="text-blue-600 hover:text-blue-800 font-medium">
              Go to League →
            </Link>
          </div>
          
          <div className="bg-green-50 p-4 rounded-lg border border-green-100">
            <h2 className="text-xl font-semibold mb-2 text-green-800">Predictions</h2>
            <p className="mb-4">Get match predictions and championship probabilities powered by machine learning.</p>
            <Link to="/predictions" className="text-green-600 hover:text-green-800 font-medium">
              View Predictions →
            </Link>
          </div>
          
          <div className="bg-purple-50 p-4 rounded-lg border border-purple-100">
            <h2 className="text-xl font-semibold mb-2 text-purple-800">Team Analysis</h2>
            <p className="mb-4">Detailed performance analysis for each team in the league.</p>
            <Link to="/team/Lions/analysis" className="text-purple-600 hover:text-purple-800 font-medium">
              View Team Analysis →
            </Link>
          </div>
        </div>
      </div>
      
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-2xl font-bold mb-4">How It Works</h2>
        <ol className="list-decimal pl-5 space-y-3">
          <li>Initialize the league with default teams (Lions, Tigers, Bears, Wolves)</li>
          <li>Play matches week by week or simulate the entire season at once</li>
          <li>View standings, match results, and team statistics</li>
          <li>Get AI-powered predictions for future matches</li>
          <li>Analyze team performance and head-to-head records</li>
        </ol>
      </div>
    </div>
  );
};

export default Home;