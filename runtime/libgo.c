#define NULL 0

void __go_runtime_error (int i){}

void *memcpy (void* dst, const void* src, int nSize)
{
    if (dst == NULL || src == NULL)
    {
        return NULL;
    }
    char* cDst = (char*)dst;
    char* cSrc = (char*)src;    
    while(nSize-->0){
        *cDst++ = *cSrc++;
    }
    return dst;
}
int memcmp(const void *buffer1,const void *buffer2,int count)
{
   if (!count)
      return(0);
   while ( --count && *(char *)buffer1 == *(char *)buffer2)
   {
      buffer1 = (char *)buffer1 + 1;
        buffer2 = (char *)buffer2 + 1;
   }
   return( *((unsigned char *)buffer1) - *((unsigned char *)buffer2) );
}
void	runtime_panicstring(const char* buf)
{}
