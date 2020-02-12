import ButtonShowMore from 'components/button_show_more'
import EpisodePreview from 'components/episode_preview'
import { Episode } from 'models'

export interface StateToProps {
  feed: Episode[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (offset: number) => void
}

const SubscriptionsFeed: React.SFC<StateToProps & DispatchToProps> = ({
  feed,
  receivedAll,
  isLoadingMore,
  loadMore,
}) => {
  if (feed.length === 0) {
    return (
      <div className="mt-8">
        <div className="text-2xl text-gray-900 tracking-wide">
          {'Latest episodes from your subscriptions show up here.'}
        </div>
        <div className="mt-2 text-default text-gray-800 tracking-wide">
          {'Subscribe to your favourite podcasts to get started.'}
        </div>
      </div>
    )
  }

  return (
    <div>
      <h1 className="pt-3 pb-1 text-xl tracking-wide text-gray-800 font-medium">
        {'Feed'}
      </h1>
      <hr className="mb-3" />
      {feed.map((episode) => (
        <div key={episode.id} className="mb-6">
          <EpisodePreview episodeId={episode.id} />
        </div>
      ))}
      {!receivedAll && (
        <div className="w-full h-10 mx-auto my-6">
          <ButtonShowMore
            isLoading={isLoadingMore}
            loadMore={() => loadMore(feed.length)}
          />
        </div>
      )}
    </div>
  )
}

export default SubscriptionsFeed
