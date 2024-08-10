import React, { useState, useEffect } from 'react';
import { Github, Linkedin } from 'lucide-react';
import Login from './Login';
import Dashboard from './Dashboard';
import { isAuthenticated } from '../services/authService';

interface ProfileProps {
  name: string;
  title: string;
  bio: string;
  skills: string[];
  github: string;
  linkedin: string;
}

const quotes = [
  "The only way to do great work is to love what you do. - Steve Jobs",
  "Code is like humor. When you have to explain it, it's bad. - Cory House",
  "First, solve the problem. Then, write the code. - John Johnson",
  "The best error message is the one that never shows up. - Thomas Fuchs",
  "It's not a bug – it's an undocumented feature. - Anonymous",
  "Talk is cheap. Show me the code. - Linus Torvalds",
  "Programming isn't about what you know; it's about what you can figure out. - Chris Pine",
  "The most disastrous thing that you can ever learn is your first programming language. - Alan Kay",
  "The best way to predict the future is to implement it. - David Heinemeier Hansson",
  "Any fool can write code that a computer can understand. Good programmers write code that humans can understand. - Martin Fowler"
];

const Profile: React.FC<ProfileProps> = ({
  name,
  title,
  bio,
  skills,
  github,
  linkedin
}) => {
  const [showLogin, setShowLogin] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [quote, setQuote] = useState("");
  const [fadeIn, setFadeIn] = useState(true);

  useEffect(() => {
    setIsLoggedIn(isAuthenticated());
    
    const changeQuote = () => {
      setFadeIn(false);
      setTimeout(() => {
        setQuote(quotes[Math.floor(Math.random() * quotes.length)]);
        setFadeIn(true);
      }, 500); // Wait for fade out, then change quote and fade in
    };

    changeQuote(); // Initial quote

    const intervalId = setInterval(changeQuote, 30000);

    return () => clearInterval(intervalId);
  }, []);

  const handleLogin = () => {
    setIsLoggedIn(true);
    setShowLogin(false);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
  };

  if (isLoggedIn) {
    return <Dashboard onLogout={handleLogout} />;
  }

  return (
    <div className="min-h-screen bg-black text-white p-8 relative overflow-hidden">
      {/* Abstract shapes */}
      <div className="absolute top-20 right-20 w-20 h-20 bg-green-300 opacity-50 rounded-full"></div>
      <div className="absolute bottom-20 left-20 w-40 h-40 bg-purple-500 opacity-50 transform rotate-45"></div>
      
      {/* Header */}
      <header className="flex justify-between items-center mb-16">
        <h1 className="text-3xl font-bold text-green-300">it's {name}</h1>
        <div className="flex space-x-4 items-center">
          <a href={github} target="_blank" rel="noopener noreferrer" className="text-green-300 hover:text-green-400">
            <Github size={24} />
          </a>
          <span className="text-gray-500 cursor-not-allowed">
            <Linkedin size={24} />
          </span>
          <button
            onClick={() => setShowLogin(true)}
            className="bg-green-300 text-black px-4 py-2 rounded hover:bg-green-400 transition-colors"
          >
            Login
          </button>
        </div>
      </header>
      
      {/* Main content */}
      {showLogin ? (
        <Login onClose={() => setShowLogin(false)} onLogin={handleLogin} />
      ) : (
        <div className="flex flex-col lg:flex-row justify-between items-start">
          <div className="lg:w-1/2 mb-8 lg:mb-0">
            <h2 className="text-6xl font-bold mb-4">{title}</h2>
            <p className="text-xl mb-8">{bio}</p>
            
            <div className="mb-8">
              <h3 className="text-green-300 mb-4 text-2xl">Quote of the day:</h3>
              <div className="bg-gray-800 p-4 rounded-lg h-32 flex items-center">
                <p className={`text-green-300 italic transition-opacity duration-500 ${fadeIn ? 'opacity-100' : 'opacity-0'}`}>
                  "{quote}"
                </p>
              </div>
            </div>
          </div>
          
          <div className="lg:w-1/2 lg:pl-8">
            <div className="mb-8">
              <div className="relative">
                <div className="w-64 h-64 bg-green-300 absolute top-4 left-4 rounded-lg"></div>
                <img 
                  src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/3f0f0d8c-eacc-48db-abd0-e7c953bbc059/dgp9kxr-cfe543ed-dec8-47cb-9f60-7684a7961ace.png/v1/fill/w_1081,h_739,q_70,strp/shadowed_cyborg__radiant_blue_aura_by_abdallahalswaiti_dgp9kxr-pre.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7ImhlaWdodCI6Ijw9MTMxNCIsInBhdGgiOiJcL2ZcLzNmMGYwZDhjLWVhY2MtNDhkYi1hYmQwLWU3Yzk1M2JiYzA1OVwvZGdwOWt4ci1jZmU1NDNlZC1kZWM4LTQ3Y2ItOWY2MC03Njg0YTc5NjFhY2UucG5nIiwid2lkdGgiOiI8PTE5MjAifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6aW1hZ2Uub3BlcmF0aW9ucyJdfQ.t43JxJW6JHlUaQJ5flM7GvS-WrY5ttCCc04SRVU4EfE" 
                  alt="Profile" 
                  className="w-64 h-64 object-cover relative z-10 rounded-lg"
                />
              </div>
            </div>
            <h3 className="text-green-300 mb-4 text-2xl">Skills:</h3>
            <ul className="list-disc list-inside">
              {skills.map((skill, index) => (
                <li key={index} className="text-white text-lg mb-2">{skill}</li>
              ))}
            </ul>
          </div>
        </div>
      )}
    </div>
  );
};

export default Profile;