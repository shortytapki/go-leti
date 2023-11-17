async function getRhymes() {
  const response = await fetch("/api/rhymes", { headers: { 'Content-Type': 'application/json' } });
  const jsonResponse = await response.json();
  return jsonResponse;
}

async function getSongs() {
  const response = await fetch("/api/songs", { headers: { 'Content-Type': 'application/json' } });
  const jsonResponse = await response.json();
  return jsonResponse;
}


document.addEventListener('DOMContentLoaded', async () => {
  const rhymesContainer = document.getElementById("rhymes");
  const songsContainer = document.getElementById("songs");

  const rhymesForm = document.getElementById('rhymesForm');
  const songsForm = document.getElementById('songsForm');

  rhymesForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const data = new FormData(e.target);
    const rhyme = data.get('rhyme');
    rhymesContainer.insertAdjacentHTML('beforeend', `<p class="text-lg">${rhyme}</p>`);
    fetch('/api/rhymes', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ Text: rhyme }) })
  });

  songsForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const data = new FormData(e.target);
    const artist = data.get('artist');
    const title = data.get('title');
    console.log(title, artist)
    songsContainer.insertAdjacentHTML('beforeend', `<p class="text-lg">${artist}-${title}</p>`);
    fetch('/api/songs', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ Artist: artist, Title: title }) })
  });


  const rhymes = await getRhymes();
  const rhymesHtml = rhymes.reduce((acc, rhyme) => acc += `<p class="text-lg">${rhyme.Text}</p>`, '')
  rhymesContainer.insertAdjacentHTML('beforeend', rhymesHtml);

  const songs = await getSongs();
  const songsHtml = songs.reduce((acc, song) => acc += `<p class="text-lg">${song.Artist}-${song.Title}</p>`, '')
  songsContainer.insertAdjacentHTML('beforeend', songsHtml);

})
