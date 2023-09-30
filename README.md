<h1 align="center">Blockchain-Based eVault for Legal Records</h1>


---
<h2>Project Structure</h2>

```
$PROJECT_ROOT
├──client/
│    ├── public/           # public static assets 
│    └── src/
│        ├── assets/       # Media files
│        ├── components/   # Reusable UI components
│        ├── utils/        # Custom hooks and helpers
│        ├── views/        # Linked pages
│        ├── routes/       # Route configurations
│        ├── styles/       # CSS styles
│        ├── App.Vue       # Main application component
│        └── index.js      # Entry point for rendering the app
│
├── server/
│    ├── app/
│    │    └── main.go      # server entry point
│    ├── db/
│    ├── models/
│    ├── controllers/
│    ├── middlewares/
│    ├── routes/
│    └── go.mod
│
│
├── ethereum/
│    ├── contracts/
│    ├── migrations/
│    └── test/
│
├── doc/                # Assects for project documentation
└── README.md           # Project documentation
```

---