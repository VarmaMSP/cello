import Grid from 'components/grid'
import { imageUrl } from 'components/utils'
import Link from 'next/link'
import { Podcast } from 'types/app'

export interface StateToProps {
  subscriptions: Podcast[]
}

const SubscriptionsList: React.SFC<StateToProps> = ({ subscriptions }) => {
  return (
    <>
      <h2 className="text-xl text-gray-900 pb-8">{'Your subscriptions'}</h2>
      <Grid
        cols={{ LG: 7, MD: 6, SM: 3 }}
        totalRowSpacing={{ LG: 12, MD: 10, SM: 10 }}
        className="md:mb-4 mb-2"
      >
        {subscriptions.map((podcast) => (
          <Link
            href={{ pathname: '/podcasts', query: { podcastId: podcast.id } }}
            as={`/podcasts/${podcast.id}`}
            key={podcast.id}
          >
            <a>
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border cursor-pointer"
                src={imageUrl(podcast.id, 'md')}
              />
            </a>
          </Link>
        ))}
      </Grid>
    </>
  )
}

export default SubscriptionsList
