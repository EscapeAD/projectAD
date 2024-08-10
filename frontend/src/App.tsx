import React from 'react';
import Profile from './components/Profile';

const App: React.FC = () => {
  const profileData = {
    name: "Adam",
    title: "Full Stack Developer",
    bio: "I craft robust and scalable full stack solutions with exceptional user experiences.",
    skills: [
      "Full Stack Development",
      "Java",
      "Spring Boot",
      "DevOps",
      "React Front End"
    ],
    github: "https://github.com/escapead",
    linkedin: "#"  // This is now just a placeholder
  };

  return (
    <div className="App">
      <Profile {...profileData} />
    </div>
  );
};

export default App;