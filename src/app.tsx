import { WorldMap } from 'components/world-map';
import { Navbar } from 'components/navbar';

const App = () => {
  return (
    <div className="app">
      <Navbar />
      <header className="app-header">Map</header>
      <WorldMap />
    </div>
  );
};

export default App;
