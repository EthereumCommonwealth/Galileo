import React from 'react'
import Link from 'gatsby-link'

const SecondPage = () => (
  <div className="container">
    <h1>Hi from the second page</h1>
    <p>Welcome to page 2</p>
    <br />
    <Link
      className="btn"
      to="/"
    >
      Go back to the homepage
    </Link>
  </div>
)

export default SecondPage
