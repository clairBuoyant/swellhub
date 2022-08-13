import { WorldMap } from '@components/world-map';
import { Navbar } from '@components/navbar';

export default function App() {
  return (
    <div className="app">
      <Navbar />
      <header className="app-header">Map</header>
      <WorldMap />
    </div>
  );
}
