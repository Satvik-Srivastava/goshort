import { useState } from 'react';
import './index.css';
import URLShortener from './components/URLShortener';

function App() {
  return (
    <div className="min-h-screen bg-white">
      <URLShortener />
    </div>
  );
}

export default App;
