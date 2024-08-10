import React, { useState, useEffect } from 'react';
import { Github, Linkedin, Twitter } from 'lucide-react';
import Login from './Login';
import Dashboard from './Dashboard';
import { isAuthenticated } from '../services/authService';

interface Skill {
  name: string;
  level: string;
}

interface ProfileProps {
  name: string;
  title: string;
  bio: string;
  age: number;
  experience: number;
  email: string;
  skills: Skill[];
  github: string;
  linkedin: string;
  twitter: string;
}

const Profile: React.FC<ProfileProps> = ({
  name,
  title,
  bio,
  age,
  experience,
  email,
  skills,
  github,
  linkedin,
  twitter
}) => {
  const [showLogin, setShowLogin] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    setIsLoggedIn(isAuthenticated());
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
          <a href={linkedin} target="_blank" rel="noopener noreferrer" className="text-green-300 hover:text-green-400">
            <Linkedin size={24} />
          </a>
          <a href={twitter} target="_blank" rel="noopener noreferrer" className="text-green-300 hover:text-green-400">
            <Twitter size={24} />
          </a>
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
        <div className="flex flex-col lg:flex-row justify-between items-center">
          <div className="lg:w-1/2 mb-8 lg:mb-0">
            <h2 className="text-6xl font-bold mb-4">{title}</h2>
            <p className="text-xl mb-8">{bio}</p>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
              <div>
                <h3 className="text-green-300 mb-2">Age: {age}</h3>
                <h3 className="text-green-300 mb-2">Experience: {experience} years</h3>
                <h3 className="text-green-300 mb-2">Email: {email}</h3>
              </div>
              <div>
                <h3 className="text-green-300 mb-2">Skills:</h3>
                <ul>
                  {skills.map((skill, index) => (
                    <li key={index} className="text-white">{skill.name} - {skill.level}</li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
          
          <div className="lg:w-1/2 relative">
            <div className="w-64 h-64 bg-green-300 absolute top-4 left-4 rounded-lg"></div>
            <img 
              src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/3f0f0d8c-eacc-48db-abd0-e7c953bbc059/dgp9kxr-cfe543ed-dec8-47cb-9f60-7684a7961ace.png/v1/fill/w_1081,h_739,q_70,strp/shadowed_cyborg__radiant_blue_aura_by_abdallahalswaiti_dgp9kxr-pre.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7ImhlaWdodCI6Ijw9MTMxNCIsInBhdGgiOiJcL2ZcLzNmMGYwZDhjLWVhY2MtNDhkYi1hYmQwLWU3Yzk1M2JiYzA1OVwvZGdwOWt4ci1jZmU1NDNlZC1kZWM4LTQ3Y2ItOWY2MC03Njg0YTc5NjFhY2UucG5nIiwid2lkdGgiOiI8PTE5MjAifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6aW1hZ2Uub3BlcmF0aW9ucyJdfQ.t43JxJW6JHlUaQJ5flM7GvS-WrY5ttCCc04SRVU4EfE" 
              alt="Profile" 
              className="w-64 h-64 object-cover relative z-10 rounded-lg"
            />
          </div>
        </div>
      )}
    </div>
  );
};

export default Profile;