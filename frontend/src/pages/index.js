import React from 'react'
import Link from 'gatsby-link'
import SampleComponent from '../components/SampleComponent.jsx'

const IndexPage = () => (
  <div className="container">
    <h1>Hi people</h1>
    <p>Welcome Callisto Gatsby Starter</p>
    <p>Now go build something great.</p>
    {/* <SampleComponent /> */}
    <br />
    <Link
      className="btn"
      to="/page-2/"
    >
      Go to page 2
    </Link>
  </div>
)

export default IndexPage
