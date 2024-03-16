import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import LoginForm from './components/LoginForm';
import SignUpForm from './components/Signup';
import ForgotPasswordForm from './components/foregtPasswordform';
import FileUploadForm from '../src/components/FileUploadForm';
import MapWithMarkers from '../src/components/MapWithMarkers';
import './App.css';

function App() {
  const [uploadedFiles, setUploadedFiles] = useState([]);
  const [selectedFile, setSelectedFile] = useState(null);

  const handleFileChange = (e) => {
    setSelectedFile(e.target.files[0]);
  };

  const handleFileUpload = () => {
    const formData = new FormData();
    formData.append('file', selectedFile);

    fetch('http://localhost:8080', {
      method: 'POST',
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => {
        setUploadedFiles([...uploadedFiles, data]); // Assuming the response contains file details
        setSelectedFile(null);
      })
      .catch((error) => {
        console.error('Error uploading file:', error);
      });
  };

  return (
    <Router>
      <div>
        <Navbar />
        <div className="container">
          <h1>Geospatial Data Management</h1>
          <Routes>
            <Route path="/login" element={<LoginForm />} />
            <Route path="/signup" element={<SignUpForm />} />
            <Route path="/forgot-password" element={<ForgotPasswordForm />} />
            <Route path="/upload" element={<FileUploadForm onFileChange={handleFileChange} onFileUpload={handleFileUpload} selectedFile={selectedFile} />} />
            <Route path="/map" element={<MapWithMarkers uploadedFiles={uploadedFiles} />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
