document.getElementById('readAll').addEventListener('click', async () => {
  const res = await fetch('/api/read/all')
  const data = await res.json()
  document.getElementById('output').textContent = JSON.stringify(data, null, 2)
})