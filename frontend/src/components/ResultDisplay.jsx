import { useState } from 'react';

const ResultDisplay = ({ result }) => {
  const [copied, setCopied] = useState(false);
  const [showPreview, setShowPreview] = useState(false);

  const handleCopy = () => {
    navigator.clipboard.writeText(result.custom_short);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const handleRedirect = () => {
    window.open(result.custom_short, '_blank');
  };

  return (
    <div className="mt-10 slide-up">
      <div className="p-6 border-2 border-black rounded-xl  text-white shadow-lg hover:shadow-xl transition-all duration-300">
        <div className="flex items-center gap-2 mb-4">
          <span className="text-2xl">🎉</span>
          <h2 className="text-xl font-bold">Your Short URL</h2>
        </div>

        {/* Short URL Display */}
        <div className="mb-5">
          <p className="text-xs font-semibold text-gray-300 mb-2 tracking-wide">SHORT URL</p>
          <div className="flex items-center gap-3 bg-gray-900 p-4 rounded-lg">
            <input
              type="text"
              value={result.custom_short}
              readOnly
              className="flex-1 bg-transparent text-white font-mono text-sm focus:outline-none"
            />
            <button
              onClick={handleCopy}
              className="px-4 py-2 bg-white text-black font-bold rounded-lg hover:bg-gray-100 transition-all duration-200 text-sm whitespace-nowrap"
            >
              {copied ? '✓ Copied' : 'Copy'}
            </button>
          </div>
        </div>

        {/* Original URL Display */}
        <div className="mb-5">
          <p className="text-xs font-semibold text-gray-300 mb-2 tracking-wide">ORIGINAL URL</p>
          <p className="text-xs text-gray-400 break-all line-clamp-2 hover:line-clamp-none transition-all cursor-pointer hover:text-gray-300">
            {result.url}
          </p>
        </div>

        {/* Rate Limit Info */}
        <div className="mb-5 grid grid-cols-2 gap-4">
          <div className="p-3 bg-gray-900 rounded-lg">
            <p className="text-xs text-gray-400 font-semibold mb-1">REQUESTS LEFT</p>
            <p className="text-xl font-black text-white">{result.rate_limit}</p>
          </div>
          <div className="p-3 bg-gray-900 rounded-lg">
            <p className="text-xs text-gray-400 font-semibold mb-1">RESET IN</p>
            <p className="text-xl font-black text-white">{result.rate_limit_reset}m</p>
          </div>
        </div>

        {/* Action Buttons */}
        <div className="flex gap-3">
          <button
            onClick={handleRedirect}
            className="flex-1 py-3 bg-white text-black font-bold rounded-lg hover:bg-gray-100 transition-all duration-200 flex items-center justify-center gap-2"
          >
            <span>🔗</span> Open Link
          </button>
        </div>
      </div>
    </div>
  );
};

export default ResultDisplay;
