import React from 'react';
import { Link } from 'react-router-dom';

const Header: React.FC = () => {
  return (
    <header className="bg-blue-800 text-white shadow-lg">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <Link to="/" className="text-2xl font-bold hover:text-blue-200 transition-colors">
          LeagueSimulator
        </Link>
        <nav>
          <ul className="flex space-x-6">
            <li>
              <Link to="/" className="hover:text-blue-200 transition-colors">Home</Link>
            </li>
            <li>
              <Link to="/league" className="hover:text-blue-200 transition-colors">League</Link>
            </li>
            <li>
              <Link to="/predictions" className="hover:text-blue-200 transition-colors">Predictions</Link>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  );
};

export default Header;