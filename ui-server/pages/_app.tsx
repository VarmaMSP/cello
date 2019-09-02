import '../styles/index.css'
import NavBarTop from '../components/navbar_top'
import AudioPlayer from '../components/audio_player'
import NavbarSide from '../components/navbar_side/navbar_side'
import MainContent from '../components/main_content'
import { makeStore } from '../store'
import { AppContext } from 'types/next'

import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { Container, AppProps } from 'next/app'
import withRedux from 'next-redux-wrapper'

export default withRedux(makeStore)(
  class MyApp extends Component<AppProps> {
    static async getInitialProps({ Component, ctx }: AppContext) {
      let pageProps = {}
      if (Component.getInitialProps) {
        pageProps = await Component.getInitialProps(ctx)
      }
      return { pageProps }
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
                'https://raw.githubusercontent.com/anars/blank-audio/master/15-minutes-of-silence.mp3'
              }
              audioType={''}
            />
          </Provider>
        </Container>
      )
    }
  },
)
