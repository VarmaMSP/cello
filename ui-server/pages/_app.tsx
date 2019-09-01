import '../styles/index.css'
import NavBarTop from '../components/navbar_top'
import AudioPlayer from '../components/audio_player'
import NavbarSide from '../components/navbar_side/navbar_side'
import MainContent from '../components/main_content'
import { makeStore, AppState } from '../store'
import { AppActions } from '../types/actions'

import React from 'react'
import { NextPageContext } from 'next'
import App, { Container, AppContext } from 'next/app'
import withRedux, { NextJSContext } from 'next-redux-wrapper'
import { Provider } from 'react-redux'

export default withRedux(makeStore)(
  class MyApp extends App {
    static async getInitialProps(c: AppContext) {
      c.ctx as NextPageContext & NextJSContext<AppState, AppActions>
      return { pageProps: {} }
    }

    render() {
      const { Component, pageProps } = this.props
      return (
        <Container>
          <Provider store={(this.props as any).store}>
            <NavBarTop />
            <NavbarSide />
            <MainContent>
              <Component {...pageProps} />
            </MainContent>
            <AudioPlayer
              podcast={'WTF M. Night Shyamalan!?!?!?!'}
              episode={'Part 3: Critics, Trolls and Narfs'}
              podcastId={''}
              episodeId={''}
              albumArt={
                'https://is5-ssl.mzstatic.com/image/thumb/Podcasts123/v4/bf/cb/94/bfcb9429-69f8-6b4a-e639-510b4bbe25a5/mza_7508403085647585170.jpg/400x400.jpg'
              }
              audioSrc={
                'http://traffic.libsyn.com/joeroganexp/p1338.mp3?dest-id=19997'
              }
              audioType={''}
            />
          </Provider>
        </Container>
      )
    }
  },
)
