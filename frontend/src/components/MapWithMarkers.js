// MapWithMarkers.js
import React from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import '../CSS/mapmarker.css'; // Make sure to adjust the path as needed

const MapWithMarkers = ({ uploadedFiles }) => {
  return (
    <div className="map-container">
      <h2 className="map-title">Map</h2>
      <MapContainer center={[51.505, -0.09]} zoom={13} className="map">
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
        {uploadedFiles.map((file, index) => (
          <Marker key={index} position={[file.latitude, file.longitude]}>
            <Popup>
              <strong className="popup-content">{file.name}</strong>
            </Popup>
          </Marker>
        ))}
      </MapContainer>
    </div>
  );
};

export default MapWithMarkers;
