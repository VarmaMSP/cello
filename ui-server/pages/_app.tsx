import React from 'react'
import App, {Container} from 'next/app'

import '../styles/index.css'
import NavBarTop from '../components/navbar_top';

export default class MyApp extends App {
  static async getInitialProps ({ Component, ctx }) {
    let pageProps = {}
    if (Component.getInitialProps) {
      pageProps = await Component.getInitialProps(ctx)
    }
    return {pageProps}
  }

  render () {
    const {Component, pageProps} = this.props
    return <Container>
      <NavBarTop showFullSearchBar={true} toggleFullSearchBar={() => {}}/>
      <Component {...pageProps} />
    </Container>
  }
}
