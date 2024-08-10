import React from 'react';
import Profile from './components/Profile';

const App: React.FC = () => {
  const profileData = {
    name: "Adam",
    title: "Full Stack Developer",
    bio: "I craft robust and scalable full stack solutions with exceptional user experiences.",
    age: 28,
    experience: 5,
    email: "adam@example.com",
    skills: [
      { name: "React", level: "Expert" },
      { name: "Node.js", level: "Advanced" },
      { name: "TypeScript", level: "Intermediate" },
      { name: "Python", level: "Advanced" },
      { name: "AWS", level: "Intermediate" },
    ],
    github: "https://github.com/adamdev",
    linkedin: "https://linkedin.com/in/adamdev",
    twitter: "https://twitter.com/adamdev"
  };

  return (
    <div className="App">
      <Profile {...profileData} />
    </div>
  );
};

export default App;