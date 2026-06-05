import { useState } from 'react';
import ResultDisplay from './ResultDisplay';

const URLShortener = () => {
  const [url, setUrl] = useState('');
  const [customShort, setCustomShort] = useState('');
  const [expiry, setExpiry] = useState('24');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState(false);

  const handleShorten = async (e) => {
    e.preventDefault();
    setError('');
    setResult(null);
    setSuccess(false);
    setLoading(true);

    if (!url.trim()) {
      setError('Please enter a URL');
      setLoading(false);
      return;
    }

    try {
      const payload = {
        url: url.trim(),
        custom_short: customShort.trim(),
        expiry: parseInt(expiry) || 24,
      };

      const apiUrl = '/api/v1/shorten';
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to shorten URL');
      }

      const data = await response.json();
      setResult(data);
      setSuccess(true);
      setUrl('');
      setCustomShort('');
      setTimeout(() => setSuccess(false), 3000);
    } catch (err) {
      setError(err.message || 'An error occurred while shortening the URL');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-white via-gray-50 to-white flex items-center justify-center p-4">
      <div className="w-full max-w-md fade-in">
        <div className="mb-12 text-center">
          <div className="inline-block mb-4 p-3 bg-black rounded-full transform hover:scale-110 transition-transform duration-300">
            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.658 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
            </svg>
          </div>
          <h1 className="text-5xl font-black text-black mb-2">GoShort</h1>
          <p className="text-gray-500 text-lg">Create short, memorable links in seconds</p>
        </div>

        <form onSubmit={handleShorten} className="space-y-5">
          <div className="slide-up" style={{animationDelay: '0.1s'}}>
            <label className="block text-sm font-semibold text-black mb-3 tracking-wide">
              📎 Your URL
            </label>
            <input
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="https://example.com/very/long/url"
              className="w-full px-5 py-3 border-2 border-gray-300 rounded-xl focus:border-black focus:ring-2 focus:ring-black focus:ring-opacity-10 bg-white text-black placeholder-gray-400 transition-all duration-200 shadow-sm hover:shadow-md"
            />
          </div>

          <div className="slide-up" style={{animationDelay: '0.2s'}}>
            <label className="block text-sm font-semibold text-black mb-3 tracking-wide">
              ✨ Custom Short (optional)
            </label>
            <input
              type="text"
              value={customShort}
              onChange={(e) => setCustomShort(e.target.value)}
              placeholder="my-awesome-link"
              className="w-full px-5 py-3 border-2 border-gray-300 rounded-xl focus:border-black focus:ring-2 focus:ring-black focus:ring-opacity-10 bg-white text-black placeholder-gray-400 transition-all duration-200 shadow-sm hover:shadow-md"
            />
          </div>

          <div className="slide-up" style={{animationDelay: '0.3s'}}>
            <label className="block text-sm font-semibold text-black mb-3 tracking-wide">
              ⏰ Link Expiry
            </label>
            <select
              value={expiry}
              onChange={(e) => setExpiry(e.target.value)}
              className="w-full px-5 py-3 border-2 border-gray-300 rounded-xl focus:border-black focus:ring-2 focus:ring-black focus:ring-opacity-10 bg-white text-black transition-all duration-200 shadow-sm hover:shadow-md font-medium"
            >
              <option value="1">1 hour</option>
              <option value="24">24 hours</option>
              <option value="48">48 hours</option>
              <option value="168">1 week</option>
              <option value="720">1 month</option>
            </select>
          </div>

          {success && (
            <div className="p-4 bg-green-50 border-2 border-green-300 rounded-xl slide-up">
              <p className="text-green-700 text-sm font-semibold flex items-center gap-2">
                <span className="text-lg">✓</span> URL shortened successfully!
              </p>
            </div>
          )}

          {error && (
            <div className="p-4 bg-red-50 border-2 border-red-300 rounded-xl slide-up">
              <p className="text-red-700 text-sm font-semibold flex items-center gap-2">
                <span className="text-lg">✕</span> {error}
              </p>
            </div>
          )}

          <button
            type="submit"
            disabled={loading}
            className="slide-up w-full py-3 bg-black text-white font-bold rounded-xl hover:bg-gray-900 disabled:bg-gray-400 disabled:cursor-not-allowed transition-all duration-200 shadow-lg hover:shadow-xl text-lg tracking-wide"
            style={{animationDelay: '0.4s'}}
          >
            {loading ? (
              <span className="flex items-center justify-center gap-2">
                <span className="inline-block w-2 h-2 bg-white rounded-full animate-bounce"></span>
                Processing...
              </span>
            ) : (
              '🚀 Shorten URL'
            )}
          </button>
        </form>

        {result && <ResultDisplay result={result} />}
      </div>
    </div>
  );
};

export default URLShortener;
