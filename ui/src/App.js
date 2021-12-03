import FoundationList from './FoundationList';

function SearchBar() {
  return (
    <>
      <form>
        <input class="border-purple-400" type="text" placeholder="Search..." />
        <input type="submit" value="Search" />
      </form>
    </>
  )
}

function App() {
  const [foundations, setFoundations] = useState([])

  return (
    <div>
      <h1>Foundation Match</h1>
      <SearchBar />
      {foundations && <FoundationList list={foundations}/>}
    </div>
  );
}

export default App;
