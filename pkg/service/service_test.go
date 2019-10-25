package service_test

import (
	"context"
	"errors"
	"fmt"

	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
	"github.com/penkovski/ginkgo-example/pkg/service"
	"github.com/penkovski/ginkgo-example/pkg/service/servicefakes"
)

var _ = Describe("Service Public API", func() {

	var (
		users    *servicefakes.FakeUsers
		storage  *servicefakes.FakeStorage
		download *servicefakes.FakeDownload

		svc *service.Service
	)

	BeforeEach(func() {
		users = new(servicefakes.FakeUsers)
		storage = new(servicefakes.FakeStorage)
		download = new(servicefakes.FakeDownload)
		svc = service.New(users, storage, download)
	})

	Describe("Get File", func() {
		var (
			token    string
			filename string

			file *service.File
			err  error
		)

		JustBeforeEach(func() {
			file, err = svc.File(context.Background(), token, filename)
		})

		Context("when request is not authorized", func() {

			Context("but error happens while getting user", func() {
				BeforeEach(func() {
					token = ""
					users.GetReturns(nil, fmt.Errorf("internal error"))
				})
				It("should return unauthorized error", func() {
					Expect(file).To(BeNil())
					Expect(err).To(Equal(service.ErrUnauthorized))
				})
			})

			Context("and user is nil without error", func() {
				BeforeEach(func() {
					token = ""
					users.GetReturns(nil, nil)
				})
				It("should return unauthorized error", func() {
					Expect(file).To(BeNil())
					Expect(err).To(Equal(service.ErrUnauthorized))
				})
			})
		})

		Context("when request is authorized", func() {
			BeforeEach(func() {
				users.GetReturns(&service.User{GUID: "123"}, nil)
			})

			Context("but file is not found in storage", func() {

				Context("when error is internal", func() {
					BeforeEach(func() {
						storage.FileReturns(nil, errors.New("some error"))
					})
					It("should return internal error", func() {
						Expect(file).To(BeNil())
						Expect(err).To(Equal(service.ErrInternal))
					})
				})

				Context("when file is missing from storage ", func() {
					BeforeEach(func() {
						storage.FileReturns(nil, service.ErrFileNotFound)
					})

					Context("when downloading file fails", func() {
						BeforeEach(func() {
							download.DownloadReturns(nil, errors.New("some error"))
						})
						It("should return error downloading file", func() {
							Expect(file).To(BeNil())
							Expect(err).To(Equal(service.ErrFileDownload))
						})
					})

					Context("when downloading file succeeds", func() {
						BeforeEach(func() {
							download.DownloadReturns(&service.File{Name: "myfile"}, nil)
						})

						Context("but saving file in storage fails", func() {
							BeforeEach(func() {
								storage.SaveFileReturns(errors.New("some error"))
							})
							It("should return error downloading file", func() {
								Expect(file).To(BeNil())
								Expect(err).To(Equal(service.ErrFileDownload))
							})
						})

						Context("and saving file in storage succeeds", func() {
							BeforeEach(func() {
								storage.SaveFileReturns(nil)
							})
							It("should return the file", func() {
								Expect(err).To(BeNil())
								Expect(file).ToNot(BeNil())
								Expect(file.Name).To(Equal("myfile"))
							})
						})
					})
				})
			})

			Context("and file is found in storage", func() {
				BeforeEach(func() {
					storage.FileReturns(&service.File{Name: "myfile2"}, nil)
				})
				It("should return the file", func() {
					Expect(err).To(BeNil())
					Expect(file).ToNot(BeNil())
					Expect(file.Name).To(Equal("myfile2"))
				})
			})
		})
	})
})
