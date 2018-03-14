#include <CppUTest/CommandLineTestRunner.h>

/*
 * ring buffer
 */

struct RingBuffer {
	int buf[10];
	int len;
};

RingBuffer * RingBuffer_Create() 
{
	struct RingBuffer * ring = (struct RingBuffer *)malloc(sizeof(struct RingBuffer));
	ring->len = 0;
	return ring;
}

void RingBuffer_Destroy(struct RingBuffer * ring) 
{
	delete(ring);
}

void RingBuffer_Put(struct RingBuffer * ring, int val)
{
	ring->buf[0] = val;
	ring->len += 1;
}

void RingBuffer_Buffer(struct RingBuffer * ring, int **buf, int *bufLen)
{
	int * bufTmp = (int*)malloc(sizeof(int) * 10);
	for(int i=0; i<10; i++)
	{
		bufTmp[i] = ring->buf[i];
	}
	*bufLen=10;
	*buf=bufTmp;
}

int RingBuffer_Take(struct RingBuffer * ring, int num, int **buf, int *bufLen)
{
	int ret = 1;

		if (ring->len >= num)
	{
		ret = 0;
	}

	if (0 == ret)
	{
		int * bufTmp = (int*)malloc(sizeof(int) * num);
		for(int i=0; i<num; i++)
		{
			bufTmp[i] = ring->buf[i];
		}
		*bufLen=num;
		*buf=bufTmp;
		ring->len -= num;
	}

	return ret;
}


TEST_GROUP(ring_buffer)
{
	struct RingBuffer * ring;

	void setup()
	{
	       ring = RingBuffer_Create();
	}

	void teardown()
	{
		RingBuffer_Destroy(ring);
		ring = NULL;
	}
};

TEST(ring_buffer, EmptyWhenInitialized)
{
	CHECK(ring->len == 0);
	// int bufLen;
	// int *buf;
	// RingBuffer_Buffer(ring, &buf, &bufLen);
	//
	// int bufExp[10] = {0,0,0,0,0,0,0,0,0,0,};
	// for(int i=0; i<10; i++)
	// {
	// 	// printf("%d:%d-%d\n", i, bufExp[i], buf[i]);
	// 	CHECK_EQUAL(bufExp[i], buf[i]);
	// }
	// delete(buf);
}

// TEST(ring_buffer, DataCanBeInserted)
// {
// 	RingBuffer_Put(ring, 1);
//
// 	int bufLen;
// 	int *buf;
// 	RingBuffer_Buffer(ring, &buf, &bufLen);
// 	int bufExp[10] = {1,0,0,0,0,0,0,0,0,0,};
// 	for(int i=0; i<10; i++)
// 	{
// 		// printf("%d:%d-%d\n", i, bufExp[i], buf[i]);
// 		CHECK_EQUAL(bufExp[i], buf[i]);
// 	}
// 	delete(buf);
// }
//
// TEST(ring_buffer, DataCanBeRetrieved)
// {
// 	RingBuffer_Put(ring, 1);
//
// 	int bufLen;
// 	int *buf;
// 	RingBuffer_Take(ring, 1, &buf, &bufLen);
//
// 	CHECK_EQUAL(1, buf[0]);
// 	delete(buf);
// }
//
// TEST(ring_buffer, DataCannotBeRetrievedIfEmpty)
// {
// 	RingBuffer_Put(ring, 1);
//
// 	int bufLen;
// 	int *buf;
// 	int ret = RingBuffer_Take(ring, 2, &buf, &bufLen);
// 	CHECK(ret != 0);
// 	delete(buf);
// }
//
