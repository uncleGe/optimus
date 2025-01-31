import "./App.css";
import News from "./components/News";

function App() {
  return (
    <div className="container">
      <h1 className="logo">
        optimus<span className="dot">.</span>
      </h1>
      <p>
        Optimus curates telecom industry news to keep Optimum decision-makers optimally informed.
        Built with Go and React/TypeScript, hosted on GCP.
      </p>
      <News />
    </div>
  );
}

export default App;