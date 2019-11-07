import ButtonWithIcon from 'components/button_with_icon'
import Router from 'next/router'
import { useEffect, useRef } from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getCurrentUrlPath } from 'selectors/browser/urlPath'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { AppActions, PLAY_EPISODE } from 'types/actions'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
import { formatEpisodeDuration, formatEpisodePubDate } from 'utils/format'
import ModalContainer from './components/modal_container'
import Overlay from './components/overlay'

interface StateToProps {
  episode: Episode
  podcast: Podcast
  currentUrlPath: string
}

interface DispatchToProps {
  playEpisode: (episodeId: string) => void
}

interface OwnProps {
  episodeId: string
  closeModal: () => void
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ModalEpisode: React.SFC<Props> = (props) => {
  const ref = useRef(null) as React.RefObject<HTMLDivElement>
  const { episode, podcast, currentUrlPath, closeModal, playEpisode } = props

  useEffect(() => {
    if (ref.current) {
      const elems = ref.current.getElementsByTagName('a')
      for (let i = 0; i < elems.length; ++i) {
        elems[i].setAttribute('target', '_blank')
      }
    }
  })

  const handleClickPodcastTitle = (e: React.SyntheticEvent<HTMLElement>) => {
    e.preventDefault()
    if (currentUrlPath !== `/podcasts/${podcast.id}`) {
      Router.push(
        {
          pathname: '/podcasts',
          query: { podcastId: podcast.id, activeTab: 'episodes' },
        },
        `/podcasts/${podcast.id}/episodes`,
      )
      closeModal()
    }
  }

  return (
    <Overlay background="rgba(0, 0, 0, 0.75)">
      <ModalContainer handleClose={closeModal} closeUponClicking="OVERLAY">
        <div className="flex">
          <img
            className="w-24 h-24 flex-none object-contain rounded-lg cursor-pointer border"
            src={getImageUrl(episode.podcastId, 'sm')}
          />
          <div className="flex flex-col justify-between md:pl-4 pl-3">
            <div className="text-sm leading-tight cursor-default">
              <h1 className="md:text-base font-medium text-gray-800 leading-tight pb-1 line-clamp-3">
                {episode.title}
              </h1>
              <h3
                className="text-gray-700 line-clamp-2"
                onClick={handleClickPodcastTitle}
              >
                {`by ${podcast.title}`}
              </h3>
            </div>
          </div>
        </div>

        <div className="flex justify-between align-middle px-1 mt-3 mb-5">
          <div className="text-sm text-gray-700">
            {`Published on ${formatEpisodePubDate(episode.pubDate, false)}`}
            <span className="mx-2 font-extrabold">&middot;</span>
            {formatEpisodeDuration(episode.duration)}
          </div>
          <ButtonWithIcon
            className="flex-none w-6 text-gray-600 hover:text-black"
            icon="play-outline"
            onClick={() => {
              closeModal()
              playEpisode(episode.id)
            }}
          />
        </div>

        <div
          ref={ref}
          style={{ height: '21rem' }}
          className="external-html px-1 text-sm leading-sung text-gray-800 overflow-y-auto"
          dangerouslySetInnerHTML={{ __html: episode.description }}
        />
      </ModalContainer>
    </Overlay>
  )
}

function mapStateToProps(state: AppState, props: OwnProps): StateToProps {
  const episode = getEpisodeById(state, props.episodeId)
  const podcast = getPodcastById(state, episode.podcastId)

  return { episode, podcast, currentUrlPath: getCurrentUrlPath(state) }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    playEpisode: (episodeId: string) =>
      dispatch({
        type: PLAY_EPISODE,
        episodeId,
        currentTime: 0,
      }),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ModalEpisode)
