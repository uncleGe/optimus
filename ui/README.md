# Optimus UI

Optimus is a curated news aggregator designed to keep Optimum decision-makers informed about the telecom industry. This frontend is built with React and TypeScript, styled for a clean and modern experience, and hosted on GCP.

![Optimus UI Screenshot](public/optimus-ui-main.png)

## Features

- Fetches and displays curated telecom industry news
- Optimized styling for readability and branding
- Responsive design for various screen sizes

## Tech Stack

- **React** (with Vite for fast development)
- **TypeScript** for type safety
- **CSS** for styling
- **ESLint & Prettier** for code consistency

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Node.js](https://nodejs.org/) (Latest LTS recommended)
- [Yarn](https://yarnpkg.com/) or npm

### Installation

Clone the repository and navigate to the frontend directory:

```sh
cd optimus/ui
```

Install dependencies:

```sh
yarn install
# or
npm install
```

## Customization

### Changing the Favicon

Replace `public/favicon.ico` with your desired favicon file.

### Modifying Styles

All styles are defined in `src/App.css` and component-specific stylesheets.

### Environment Variables

Create a `.env` file in the root of the UI directory with the following:

```sh
VITE_API_URL=http://localhost:8080
```

### Running the Application

To start the development server:

```sh
yarn dev
# or
npm run dev
```

This will start the UI at `http://localhost:5173`.

### Building for Production

```sh
yarn build
# or
npm run build
```

This will generate a production-ready build in the `dist` directory.

## Deployment

Deployment is managed via GCP. Ensure that environment variables are correctly configured in the hosting environment.
