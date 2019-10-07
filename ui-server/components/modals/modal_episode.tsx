import ButtonWithIcon from 'components/button_with_icon'
import { useEffect, useRef } from 'react'
import { connect } from 'react-redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
import { formatEpisodeDuration, formatEpisodePubDate } from 'utils/format'
import ModalContainer from './components/modal_container'
import Overlay from './components/overlay'

interface StateToProps {
  episode: Episode
  podcast: Podcast
}

interface OwnProps {
  episodeId: string
  closeModal: () => void
}

interface Props extends StateToProps, OwnProps {}

const ModalEpisode: React.SFC<Props> = (props) => {
  const ref = useRef(null) as React.RefObject<HTMLDivElement>
  const { episode, podcast, closeModal } = props

  useEffect(() => {
    if (ref.current) {
      const elems = ref.current.getElementsByTagName('a')
      for (let i = 0; i < elems.length; ++i) {
        elems[i].setAttribute('target', '_blank')
      }
    }
  })

  return (
    <Overlay background="rgba(0, 0, 0, 0.75)">
      <ModalContainer handleClose={closeModal} closeUponClicking="OVERLAY">
        <div className="flex">
          <img
            className="w-20 h-20 flex-none object-contain rounded-lg border cursor-pointer"
            src={getImageUrl(episode.podcastId, 'md')}
          />
          <div className="flex flex-col justify-between md:pl-4 pl-3">
            <div>
              <h1 className="text-base font-medium text-gray-900 leading-tight pb-1 line-clamp-2">
                {episode.title}
              </h1>
              <h2 className="text-sm text-gray-700 line-clamp-2">
                {`by ${podcast.title}`}
              </h2>
            </div>
          </div>
        </div>

        <div className="flex justify-between align-middle px-1 mt-3 mb-5">
          <div className="text-sm text-gray-700">
            {`Published on ${formatEpisodePubDate(episode.pubDate, false)}`}
            <span className="mx-2 font-extrabold">&middot;</span>
            {`${formatEpisodeDuration(episode.duration)} long`}
          </div>
          <ButtonWithIcon
            className="flex-none w-6 text-gray-600 hover:text-black"
            icon="play-outline"
            onClick={() => {}}
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

  return { episode, podcast }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ModalEpisode,
)
