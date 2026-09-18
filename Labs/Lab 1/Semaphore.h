#ifndef SEMAPHORE_H
#define SEMAPHORE_H 
#include <mutex>
#include <condition_variable>
#include <chrono>

/**
 * @class Semaphore
 * @brief This is a simple semaphore used to control access between threads to avoid issues.
 *
 * This semaphore allows threads to wait until they are allowed
 * to continue and signal when they are finished to avoid issues that arise.
 */


class Semaphore
{
private:
    /**
    * @brief This Holds the semaphore count.
    */
    unsigned int m_uiCount; /*!< Holds the Semaphore count */
    std::mutex m_mutex;
    std::condition_variable m_condition;

public:
    /**
     * @brief This will create a starting count for the semaphore.
     * @param uiCount would be the variable that will keep the count.
     */

    Semaphore(unsigned int uiCount=0)
          : m_uiCount(uiCount) { };

    /**
    * @brief This will make the thread wait
    */
    void Wait();

    template< typename R,typename P >

    /**
     * @brief this will wait until it has permission to continue on to the next thread.
     * @param crRelTime is the variable that will be keeping time so that this thread can run smoothly
     * @return this will variable will only become true when the permission from the signal is received.
     */
    bool Wait(const std::chrono::duration<R,P>& crRelTime);

    /**
    * @brief This will give permission for the thread to continue.
    */
    void Signal();

};

#endif
