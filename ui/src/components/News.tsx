import { useEffect, useState, useRef } from "react";

interface Article {
  title: string;
  description: string;
  url: string;
}

function News() {
  const [articles, setArticles] = useState<Article[]>([]);
  const [error, setError] = useState<string | null>(null);
  const hasFetched = useRef(false); // Keeps state across re-renders

  useEffect(() => {
    if (hasFetched.current) return; // Prevents duplicate fetches on navigation

    const fetchNews = async () => {
      try {
        const response = await fetch("http://localhost:8080/news"); // TODO: Put this string in an environment variable
        if (!response.ok) {
          throw new Error("Failed to fetch news");
        }
        const data = await response.json();
        setArticles(data);
        hasFetched.current = true; // Marks as fetched to prevent re-fetching
      } catch (err) {
        console.error("Fetch error:", err);
        setError("Failed to fetch news");
      }
    };

    fetchNews();
  }, []); // Only runs on first mount

  return (
    <div>
      <h2>Latest News</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <ul>
        {articles.map((article, index) => (
          <li key={index}>
            <h3>{article.title}</h3>
            <p>{article.description}</p>
            <a href={article.url} target="_blank" rel="noopener noreferrer">
              Read more
            </a>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default News;