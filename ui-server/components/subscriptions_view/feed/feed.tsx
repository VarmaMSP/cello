import ButtonShowMore from 'components/button_show_more'
import EpisodeListItem from 'components/episode_list_item'
import isToday from 'date-fns/isToday'
import isYesterday from 'date-fns/isYesterday'
import { Episode } from 'types/app'

export interface StateToProps {
  feed: Episode[]
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (publishedBefore: string) => void
  showEpisodeModal: (episodeId: string) => void
}

const Feed: React.SFC<StateToProps & DispatchToProps> = (props) => {
  const { feed, loadMore, isLoadingMore } = props

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
          <hr className="mt-2" />
          {episodes.length > 0 ? (
            <div>
              {episodes.map((episode) => (
                <EpisodeListItem key={episode.id} episodeId={episode.id} />
              ))}
            </div>
          ) : (
            <p className="my-6 text-gray-600 text-sm tracking-wide">
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
