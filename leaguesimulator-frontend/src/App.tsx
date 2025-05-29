import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import Home from './pages/Home';
import LeaguePage from './pages/LeaguePage';
import PredictionsPage from './pages/PredictionsPage';

const App: React.FC = () => {
  return (
    <Router>
      <div className="min-h-screen bg-gray-100">
        <Header />
        <div className="container mx-auto px-4 py-8">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/league" element={<LeaguePage />} />
            <Route path="/predictions" element={<PredictionsPage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default App;