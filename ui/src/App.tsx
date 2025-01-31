import "./App.css";
import News from "./components/News";

function App() {
  return (
    <div className="container">
      <h1 className="logo">
        optimus<span className="dot">.</span>
      </h1>
      <p className="tagline">
        Optimus curates telecom industry news to keep Optimum decision-makers optimally informed.
      </p>
      <p className="subtagline">
        (Built with React/TypeScript and Go, hosted on GCP.)
      </p>
      <News />
    </div>
  );
}

export default App;