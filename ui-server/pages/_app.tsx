import AudioPlayer from 'components/audio_player'
import MainContainer from 'components/main_container'
import NavbarSide from 'components/navbar_side/navbar_side'
import NavBarTop from 'components/navbar_top'
import withRedux from 'next-redux-wrapper'
import { AppProps, Container } from 'next/app'
import Head from 'next/head'
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { makeStore } from 'store'
import { AppContext, PageContext } from 'types/utilities'
import '../styles/index.css'

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
            <MainContainer>
              <Component {...pageProps} />
            </MainContainer>
            <AudioPlayer />
          </Provider>
        </Container>
      )
    }
  },
)
