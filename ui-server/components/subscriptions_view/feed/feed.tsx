import ButtonPlay from 'components/button_play'
import ButtonShowMore from 'components/button_show_more'
import EpisodeMeta from 'components/episode_meta'
import isToday from 'date-fns/isToday'
import isYesterday from 'date-fns/isYesterday'
import striptags from 'striptags'
import { Episode } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  feed: Episode[]
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (publishedBefore: string) => void
  showEpisodeModal: (episodeId: string) => void
}

const Feed: React.SFC<StateToProps & DispatchToProps> = (props) => {
  const { feed, loadMore, isLoadingMore, showEpisodeModal } = props

  const feedList: { title: string; episodes: Episode[] }[] = [
    { title: 'Today', episodes: [] },
    { title: 'Yesterday', episodes: [] },
    { title: 'Earlier', episodes: [] },
  ]

  for (let i = 0; i < feed.length; ++i) {
    const episode = feed[i]
    const pubDate = new Date(`${episode.pubDate} +0000`)

    if (isToday(pubDate)) {
      feedList[0].episodes.push(episode)
      continue
    }
    if (isYesterday(pubDate)) {
      feedList[1].episodes.push(episode)
      continue
    }
    feedList[2].episodes.push(episode)
  }

  return (
    <>
      {feedList.map(({ title, episodes }) => (
        <div key={title}>
          <h1 className="text-xl  text-gray-900 mt-5">{title}</h1>
          <hr className="mt-2 mb-3" />
          {episodes.length > 0 ? (
            <div>
              {episodes.map((episode) => (
                <div key={episode.id} className="flex my-5 py-2">
                  <img
                    className="w-24 h-24 mr-2 flex-none object-contain rounded-lg border cursor-default"
                    src={getImageUrl(episode.podcastId, 'md')}
                    onClick={() => showEpisodeModal(episode.id)}
                  />
                  <div className="flex-auto flex flex-col justify-between pl-3">
                    <div className="flex-auto">
                      <h1
                        className="md:text-base text-sm line-clamp-2 cursor-default"
                        onClick={() => showEpisodeModal(episode.id)}
                      >
                        {episode.title}
                      </h1>
                      <div className="mt-2 mb-2">
                        <EpisodeMeta episodeId={episode.id} />
                      </div>
                      <p
                        className="text-xs text-gray-700 leading-snug tracking-wide line-clamp-2"
                        style={{ hyphens: 'auto' }}
                      >
                        {striptags(episode.description)}
                      </p>
                    </div>
                    <div className="flex-none mt-4">
                      <ButtonPlay className="w-5" episodeId={episode.id} />
                    </div>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <p className="text-gray-600 text-sm tracking-wide">
              {'No episodes published'}
            </p>
          )}
        </div>
      ))}

      <div className="w-28 h-10 mx-auto my-6">
        <ButtonShowMore
          isLoading={isLoadingMore}
          loadMore={() =>
            feed.length > 0 && loadMore(feed[feed.length - 1].pubDate)
          }
        />
      </div>
    </>
  )
}

export default Feed
