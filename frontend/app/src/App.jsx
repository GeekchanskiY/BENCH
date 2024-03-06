import logo from './logo.svg';
import './App.css';
import ApiCaller from './components/ApiCaller';
import Jwt from './features/jwt/Jwt';

function App() {
  return (
    <div className="App">
      <header>BENCH</header>
      <ApiCaller></ApiCaller>
      <Jwt></Jwt>
    </div>
  );
}

export default App;
