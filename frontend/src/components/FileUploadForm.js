// FileUploadForm.js
import React from 'react';
import '../CSS/filefrom.css'; // Make sure to adjust the path as needed

const FileUploadForm = ({ onFileChange, onFileUpload, selectedFile }) => {
  return (
    <div className="file-upload-form-container">
      <h2 className="file-upload-form-title">Upload File</h2>
      <input type="file" className="file-input" onChange={onFileChange} />
      <button className="upload-button" onClick={onFileUpload} disabled={!selectedFile}>Upload</button>
    </div>
  );
};

export default FileUploadForm;
