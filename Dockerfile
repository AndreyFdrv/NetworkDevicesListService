FROM scratch
ADD NetworkDevicesListService /NetworkDevicesListService
ENV PORT 3000
EXPOSE 3000
CMD ["/NetworkDevicesListService"]
