import logo from './logo.svg';
import './App.css';
import ApiCaller from './components/ApiCaller';
import Jwt from './features/jwt/Jwt';
import Header from './components/Header';

function App() {
  return (
    <div className="App">
      <Header></Header>
      <ApiCaller></ApiCaller>
      <Jwt></Jwt>
      <footer></footer>
    </div>
  );
}

export default App;
