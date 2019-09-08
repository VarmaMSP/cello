import '../styles/index.css'
import NavBarTop from '../components/navbar_top'
import AudioPlayer from '../components/audio_player'
import NavbarSide from '../components/navbar_side/navbar_side'
import MainContent from '../components/main_content'
import Screen from '../components/screen'
import { makeStore } from '../store'
import { AppContext, PageContext } from 'types/utilities'

import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { Container, AppProps } from 'next/app'
import Head from 'next/head'
import withRedux from 'next-redux-wrapper'

export default withRedux(makeStore)(
  class MyApp extends Component<AppProps & PageContext> {
    static async getInitialProps({ Component, ctx }: AppContext) {
      let pageProps = {}
      if (Component.getInitialProps) {
        pageProps = await Component.getInitialProps(ctx)
      }
      return { pageProps }
    }

    render() {
      const { Component, pageProps, store } = this.props
      return (
        <Container>
          <Head>
            <title>phenopod</title>
          </Head>

          <Provider store={store}>
            <NavBarTop />
            <NavbarSide />
            <MainContent>
              <Component {...pageProps} />
            </MainContent>
            <AudioPlayer />
            <Screen />
          </Provider>
        </Container>
      )
    }
  },
)
