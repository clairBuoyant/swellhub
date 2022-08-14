import { Navbar } from '@components/navbar';
import { WorldMap } from '@components/world-map';

export default function App() {
  return (
    <div className="app">
      <Navbar />
      <header className="app-header">Map</header>
      <WorldMap />
    </div>
  );
}
