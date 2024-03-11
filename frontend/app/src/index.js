import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import store from './app/store';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Provider } from 'react-redux';
import Header from './components/Header';
import Footer from './components/Footer';
import Login from './components/Login';
import User from './components/User';
import ServiceList from './components/ServiceList';
import LogListener from './components/LogListener';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <Provider store={store}>
    
    
  <BrowserRouter>
    <Header></Header>
    <div className='content-wrap'>
        <Routes>
          <Route path="/" element={<App />}></Route>
          <Route path='/me' element={<User/>}/>
          <Route path="/login" element={<Login />}></Route>
          <Route path='/services' element={<ServiceList/>} />
          <Route path='/logs' element={<LogListener/>} />
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
