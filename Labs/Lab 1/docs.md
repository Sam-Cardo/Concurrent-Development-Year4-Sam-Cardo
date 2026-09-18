@page title Changes made to Lab 1
@tableofcontents

@section title Lab 1

This section about hello threads and what the file does
it uses two threads and a semaphore in order to control in which 
order that they run. The changes that were done to this file were:

    //Displays a message first. 

    void taskOne(std::shared_ptr<Semaphore> theSemaphore, int delay) {
        sleep(delay);

        std::cout <<"I ";
        std::cout << "must ";
        std::cout << "print ";
        std::cout << "first"<<std::endl;
        //tell taskTwo to start now
        theSemaphore->Signal();
    }

        ` // displays a message second `
            
            void taskTwo(std::shared_ptr<Semaphore> theSemaphore){
            //wait here until taskOne finishes... 
                theSemaphore->Wait();
        
            std::cout <<"This ";
            std::cout << "will ";
            sleep(5);
            std::cout << "appear ";
            std::cout << "second"<<std::endl;
        } 



        `//  displays a message that is split in to 2 sections to show how a rendezvous works `
        
        void updateTask(std::shared_ptr<Semaphore> firstSem, int numUpdates){
    
        for(int i=0;i<numUpdates;i++) {
    
            //UPDATE SHARED VARIABLE HERE!
            firstSem->Wait();
            sharedVariable++;
            firstSem->Signal();
        }
    }


