import ButtonPlay from 'components/button_play'
import ButtonShowMore from 'components/button_show_more'
import EpisodeMeta from 'components/episode_meta'
import Grid from 'components/grid'
import isToday from 'date-fns/isToday'
import isYesterday from 'date-fns/isYesterday'
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

interface Props extends StateToProps, DispatchToProps {}

const ListFeed: React.SFC<Props> = (props) => {
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
          <h1 className="text-xl text-gray-900 mt-5 mb-2">{title}</h1>
          {episodes.length > 0 ? (
            <Grid
              cols={{ LG: 2, MD: 1, SM: 1 }}
              classNameChild="flex my-2 lg:px-2 py-2 rounded-lg md:hover:bg-gray-200"
              totalRowSpacing={{ LG: 2, MD: 10, SM: 0 }}
            >
              {episodes.map((episode) => (
                <>
                  <img
                    className="w-24 h-24 flex-none object-contain rounded-lg border cursor-default"
                    src={getImageUrl(episode.podcastId, 'md')}
                    onClick={() => showEpisodeModal(episode.id)}
                  />
                  <div className="flex-auto flex flex-col justify-between pl-3">
                    <div>
                      <h1
                        className="text-sm leading-tight line-clamp-2 cursor-default"
                        onClick={() => showEpisodeModal(episode.id)}
                      >
                        {episode.title}
                      </h1>
                      <div className="mt-2">
                        <EpisodeMeta episodeId={episode.id} />
                      </div>
                    </div>
                    <ButtonPlay className="w-5" episodeId={episode.id} />
                  </div>
                </>
              ))}
            </Grid>
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
          loadMore={() => loadMore(feed[feed.length - 1].pubDate)}
        />
      </div>
    </>
  )
}

export default ListFeed
