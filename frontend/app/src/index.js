import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './components/App';
import reportWebVitals from './reportWebVitals';
import store from './app/store';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Provider } from 'react-redux';
import Header from './components/Header';
import Footer from './components/Footer';
import Login from './components/Login';
import User from './components/User';
import ServiceList from './components/services/ServiceList';
import LogListener from './components/LogListener';
import Register from './components/Register';
import Projects from './components/projects/Projects';
import CVBuilder from './components/cvBuilder/CVBuilder';
import Ressearch from './components/ressearch/Ressearch';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <Provider store={store}>
    
    
  <BrowserRouter>
    <Header></Header>
    <div className='content-wrap'>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path='/me' element={<User/>}/>
          <Route path="/login" element={<Login />} />
          <Route path='/register' element={<Register/>} />
          <Route path='/services' element={<ServiceList/>} />
          <Route path='/logs' element={<LogListener/>} />
          <Route path='/projects' element={<Projects/>} />
          <Route path='/cv' element={<CVBuilder/>} />
          <Route path='/ressearch' element={<Ressearch/>} />
        </Routes>
    </div>
    <Footer></Footer>
  </BrowserRouter>
    
    
  </Provider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
