import '../styles/main.css';
import { Link } from 'react-router-dom';

function App() {
  return <div className="index-page">
      <h3>GeekchanskiY's bench project</h3>
      <p>
        The goal of this project is to learn some new technologies,
        create full development cycle, and just to have fun while I'm on
        bench.
      </p>
      <Link to={'/services'}>Services</Link>
  </div>

}

export default App;
