import { API_URL } from '@/constants'

// slug -> buyers?page=${this.nextPage}
// slug -> buyers/${this.$route.params.id}

function fetchData (slug) {
  return fetch(`${API_URL}/${slug}`)
    .then(response => {
      if (response.ok) return response.json()
      else throw Error('Fetch data error')
    })
}

export { fetchData }
